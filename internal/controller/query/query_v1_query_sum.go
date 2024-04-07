package query

import (
	"context"
	v1 "enhanced_riskctrl/api/query/v1"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)

func (c *ControllerV1) QuerySum(ctx context.Context, req *v1.QuerySumReq) (res *v1.QuerySumRes, err error) {
	///
	if req.From == "" && req.ChainId == 0 && req.Contract == "" && req.To == "" {
		return nil, mpccode.CodeParamInvalid()
	}

	///
	cnt, err := c.enhanced_riskctrl.GetAggSum(ctx, mpcdao.QueryEnhancedRiskCtrlRes{
		From:     req.From,
		Contract: req.Contract,
		ChainId:  req.ChainId,
		StartTs:  req.StartTime,
		EndTs:    req.EndTime,
	})
	if err != nil {
		g.Log().Error(ctx, "QuerySum err:", err)
		return nil, mpccode.CodeInternalError(mpccode.TraceId(ctx))
	}
	if cnt.Cmp(bigZero) > 0 {
		endts := time.Now().Add(c.retentionDataDur)
		err = c.enhanced_riskctrl.Clear(ctx, mpcdao.QueryEnhancedRiskCtrlRes{
			From:     req.From,
			Contract: req.Contract,
			ChainId:  req.ChainId,
			StartTs:  0,
			EndTs:    endts.Unix(),
		})
		if err != nil {
			g.Log().Error(ctx, "QuerySum err:", err)
		}
	}
	return &v1.QuerySumRes{
		Result: cnt,
	}, nil
}
