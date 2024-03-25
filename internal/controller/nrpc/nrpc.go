package nrpc

import (
	"context"
	"enhanced_riskctrl/api/dataRiskCtrl"
	iquery "enhanced_riskctrl/api/query"
	v1 "enhanced_riskctrl/api/query/v1"
	"enhanced_riskctrl/internal/conf"
	"enhanced_riskctrl/internal/controller/query"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mq"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ControllerV1 struct {
	cli iquery.IQueryV1
}

func (s ControllerV1) RpcAlive(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}

func (s ControllerV1) QueryCnt(ctx context.Context, req *dataRiskCtrl.QueryReq) (*dataRiskCtrl.QueryRes, error) {
	rst, err := s.cli.QueryCnt(ctx, &v1.QueryCntReq{
		// ChainId:  req.ChainId,
	})
	if err != nil {
		return nil, err
	}
	//
	data := &dataRiskCtrl.QueryRes{
		RiskSerial: rst.Result.(string),
	}
	return data, nil
}

func Run() {
	ctx := gctx.GetInitCtx()
	nats := mq.New(conf.Config.Nrpc.NatsUrl)
	h := dataRiskCtrl.NewDataRiskCtrlHandler(ctx, nats.Conn(), new())
	var err error
	sub, err := nats.Conn().QueueSubscribe(h.Subject(), h.Subject(), h.Handler)
	if err != nil {
		panic(err)
	}
	///
	for {
		select {
		case <-ctx.Done():
			sub.Drain()
			return

		}
	}
	///
}
func new() dataRiskCtrl.DataRiskCtrlServer {
	return &ControllerV1{
		cli: query.NewV1(),
	}
}
