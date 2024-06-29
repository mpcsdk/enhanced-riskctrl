package receiver

import (
	"context"
	"encoding/json"
	"enhanced_riskctrl/internal/conf"
	"enhanced_riskctrl/internal/service"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/lib/pq"
	"github.com/mpcsdk/mpcCommon/mpcdao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
	"github.com/mpcsdk/mpcCommon/mq"
	"github.com/nats-io/nats.go/jetstream"
)

type sReceiver struct {
	ctx  context.Context
	jet  jetstream.JetStream
	cons jetstream.Consumer
	///
	enhanced_riskctrl *mpcdao.EnhancedRiskCtrl
	mpc               *mpcdao.MpcContext
}

func isDuplicateKeyErr(err error) bool {
	gerr := err.(*gerror.Error)
	if cerr, ok := gerr.Cause().(*pq.Error); ok {
		if cerr.Code == "23505" {
			return true
		}
	}
	return false
}

var zeroAddr common.Address

func new() *sReceiver {

	ctx := gctx.GetInitCtx()
	retention, err := gtime.ParseDuration(conf.Config.Cache.RetentionDataTime)
	if err != nil {
		panic(err)
	}
	///
	nats := mq.New(conf.Config.Nrpc.NatsUrl)
	jet := nats.JetStream()

	cons, err := nats.GetConsumer("enhanced_riksctrl", mq.JetStream_SyncChain, mq.JetSub_SyncChainTransfer_Latest)
	if err != nil {
		panic(err)
	}

	///
	r := g.Redis()
	_, err = r.Conn(gctx.GetInitCtx())
	if err != nil {
		panic(err)
	}
	redisAggTx := g.Redis("aggTx")
	_, err = r.Conn(gctx.GetInitCtx())
	if err != nil {
		panic(err)
	}
	///
	s := &sReceiver{
		ctx:               ctx,
		jet:               jet,
		cons:              cons,
		enhanced_riskctrl: mpcdao.NewEnhancedRiskCtrl(redisAggTx, -1),
		mpc:               mpcdao.NewMcpContet(r, conf.Config.Cache.Duration),
	}
	///
	///
	cons.Consume(func(msg jetstream.Msg) {
		tx := &entity.ChainTx{}
		json.Unmarshal(msg.Data(), tx)

		defer msg.Ack()
		g.Log().Debug(ctx, "enhancedtx:", tx)
		if tx.From == zeroAddr.String() {
			g.Log().Debug(ctx, "0 fromaddr:", tx)
			return
		}
		// filter mpcaddr tx
		ok := false
		var err error
		if ok, err = s.mpc.ExistsWalletAddr(ctx, tx.From, tx.ChainId); err != nil {
			g.Log().Error(ctx, "check mpcaddr:", tx.From, ", err:", err)
			return
		}
		////
		if !ok {
			g.Log().Debug(ctx, "check mpcaddr:", tx.From, ", not exists")
			return
		}
		g.Log().Notice(ctx, "check mpcaddr:", tx.From)
		///
		err = s.enhanced_riskctrl.InsertTx(ctx, tx)
		if err != nil {
			if isDuplicateKeyErr(err) {
				g.Log().Warning(ctx, "insertdb isDuplicateKeyErr:", tx, ", err:", err)
			} else {
				g.Log().Fatal(ctx, "insertdb:", tx, ", err:", err)
			}
		} else {
			///aggval
			err = s.enhanced_riskctrl.AggTx(ctx, tx)
			if err != nil {
				g.Log().Fatal(ctx, "agg tx:", tx, ", err:", err)
			}
		}

		g.Log().Info(ctx, "check mpcaddr record:", tx.From, tx.ChainId, tx.Height)
		//clear aggcache
		endts := time.Now().Add(retention)
		err = s.enhanced_riskctrl.Clear(ctx, mpcdao.QueryEnhancedRiskCtrlRes{
			From:     tx.From,
			Contract: tx.Contract,
			ChainId:  tx.ChainId,
			StartTs:  0,
			EndTs:    endts.Unix(),
		})
		if err != nil {
			g.Log().Error(ctx, "QuerySum err:", err)
		}
	})

	///
	return s
}

func init() {
	service.RegisterReceiver(new())
}
