package query

import (
	"context"
	"time"

	v1 "enhanced_riskctrl/api/query/v1"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)

func (c *ControllerV1) QueryCnt(ctx context.Context, req *v1.QueryCntReq) (res *v1.QueryCntRes, err error) {
	///
	if req.From == "" && req.ChainId == 0 && req.Contract == "" {
		return nil, mpccode.CodeParamInvalid()
	}

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
		return nil, mpccode.CodeInternalError(mpccode.TraceId(ctx))
	}
	if cnt > 0 {
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
	return &v1.QueryCntRes{
		Result: cnt,
	}, nil
}
