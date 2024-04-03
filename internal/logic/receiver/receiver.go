package receiver

import (
	"context"
	"encoding/json"
	"enhanced_riskctrl/internal/conf"
	"enhanced_riskctrl/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
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

func new() *sReceiver {

	ctx := gctx.GetInitCtx()
	p, err := gcmd.Parse(g.MapStrBool{
		"s,sync": false,
	})
	if err != nil {
		panic(err)
	}
	////
	if p.GetOpt("sync") == nil {
		g.Log().Notice(ctx, "no sync tx")
		return &sReceiver{}
	}
	///
	nats := mq.New(conf.Config.Nrpc.NatsUrl)
	jet, err := nats.JetStream()
	if err != nil {
		panic(err)
	}

	cons, err := nats.GetChainTxConsumer("riksctrl", mq.JetSub_ChainTx)
	if err != nil {
		panic(err)
	}

	///
	r := g.Redis("aggRiskCtrl")
	_, err = r.Conn(gctx.GetInitCtx())
	if err != nil {
		panic(err)
	}
	///
	s := &sReceiver{
		ctx:               ctx,
		jet:               jet,
		cons:              cons,
		enhanced_riskctrl: mpcdao.NewEnhancedRiskCtrl(r, conf.Config.Cache.Duration),
		mpc:               mpcdao.NewMcpContet(r, conf.Config.Cache.Duration),
	}
	///
	///
	cons.Consume(func(msg jetstream.Msg) {
		tx := &entity.ChainTx{}
		json.Unmarshal(msg.Data(), tx)
		// filter mpcaddr tx
		if ok, err := s.mpc.ExistsMpcAddr(ctx, tx.From); err != nil {
			g.Log().Error(ctx, "check mpcaddr :", tx.From, ", err:", err)
			return
		} else if !ok {
			g.Log().Info(ctx, "check mpcaddr :", tx.From, ", not exists")
			return
		}
		///
		err := s.enhanced_riskctrl.InsertTx(ctx, tx)
		if err != nil {
			g.Log().Error(ctx, "insertdb :", tx, ", err:", err)
		} else {
			g.Log().Debug(ctx, "insertdb :", tx)

			///aggval
			err := s.enhanced_riskctrl.AggTx(ctx, tx)
			if err != nil {
				g.Log().Warning(ctx, "agg tx:", tx, ", err:", err)
			}
			g.Log().Debug(ctx, "agg tx:", tx)
			//
			msg.Ack()
		}
	})

	///
	return s
}

func init() {
	service.RegisterReceiver(new())
}
