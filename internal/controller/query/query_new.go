// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package query

import (
	"enhanced_riskctrl/api/query"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)

type ControllerV1 struct{
	redis  *gredis.Redis

	enhanced_riskctrl *mpcdao.EnhancedRiskCtrl
}

func NewV1() query.IQueryV1 {
	r := g.Redis("aggRiskCtrl")
	_, err := r.Conn(gctx.GetInitCtx())
	if err != nil {
		panic(err)
	}
	return &ControllerV1{
		redis:  r,
		enhanced_riskctrl: mpcdao.NewEnhancedRiskCtrl(r),
	}
}

