package query

import (
	"context"

	v1 "enhanced_riskctrl/api/query/v1"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)

func (c *ControllerV1) Query(ctx context.Context, req *v1.QueryReq) (res *v1.QueryRes, err error) {
	g.Log().Debug(ctx, "Query req:", req)
	if req.From == "" && req.To == "" && req.Contract == "" {
		return nil, mpccode.CodeParamInvalid()
	}
	if req.StartTime >= req.EndTime {
		return nil, mpccode.CodeParamInvalid()
	}
	if req.Page < 0 || req.PageSize <= 0 {
		return nil, mpccode.CodeParamInvalid()
	}
	///
	query := &mpcdao.QueryTx{
		From: func() string {
			if req.From == "" {
				return ""
			} else {
				return common.HexToAddress(req.From).String()
			}
		}(),
		To: func() string {
			if req.To == "" {
				return ""
			} else {
				return common.HexToAddress(req.To).String()
			}
		}(),
		Contract: func() string {
			if req.Contract == "" {
				return ""
			} else {
				return common.HexToAddress(req.Contract).String()
			}
		}(),
		///
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		///
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	result, err := c.enhanced_riskctrl.Query(ctx, query)
	if err != nil {
		g.Log().Error(ctx, "Query err:", err)
		return nil, mpccode.CodeParamInvalid()
	}
	//
	res = &v1.QueryRes{}
	res.Result = result
	return res, nil
}
