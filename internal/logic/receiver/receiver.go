package receiver

import (
	"context"
	"encoding/json"
	"enhanced-riskctrl/internal/conf"
	"enhanced-riskctrl/internal/service"
	"fmt"

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
}

func AggKey(chainId int64, from, contract string) string {
	return fmt.Sprintf("%d_%s_%s", chainId, from, contract)
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

	cons, err := nats.GetConsumer("riksctrl", mq.JetSub_ChainTx)
	if err != nil {
		panic(err)
	}
	///
	s := &sReceiver{
		ctx:  ctx,
		jet:  jet,
		cons: cons,
	}
	///
	///
	cons.Consume(func(msg jetstream.Msg) {
		tx := &entity.ChainData{}
		json.Unmarshal(msg.Data(), tx)
		//based db
		mpcdao.InsertTx(ctx, tx)
		g.Log().Debug(ctx, "insertdb :", tx)
		///aggval
		err := s.aggTx(ctx, AggKey(tx.ChainId, tx.FromAddr, tx.Contract), tx)
		if err != nil {
			g.Log().Warning(ctx, "agg tx:", tx, ", err:", err)
		}
		g.Log().Debug(ctx, "agg tx:", tx)
		//
		msg.Ack()

	})

	///
	return s
}

// func (s *sReceiver) runAgg() {
// 	ticker := time.NewTicker(time.Second * 10)
// 	go func() {
// 		for {
// 			select {
// 			case <-s.ctx.Done():
// 				return
// 			case <-ticker.C:
// 				//agg data

// 				///
// 				ticker.Reset(time.Second * 10)
// 			}

// 		}
// 	}()
// }

func init() {
	service.RegisterReceiver(new())
}
