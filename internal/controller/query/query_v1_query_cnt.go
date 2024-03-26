package query

import (
	"context"

	v1 "enhanced_riskctrl/api/query/v1"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)

func (c *ControllerV1) QueryCnt(ctx context.Context, req *v1.QueryCntReq) (res *v1.QueryCntRes, err error) {
	///
	///
	cnt, err := c.enhanced_riskctrl.GetAggCnt(ctx, mpcdao.QueryEnhancedRiskCtrlRes{
		From:     req.From,
		Contract: req.Contract,
		ChainId:  req.ChainId,
		StartTs:  req.StartTime,
		EndTs:    req.EndTime,
	})
	if err != nil {
		g.Log().Error(ctx, "QueryCnt err:", err)
	}
	return &v1.QueryCntRes{
		Result: cnt,
	}, nil
}
