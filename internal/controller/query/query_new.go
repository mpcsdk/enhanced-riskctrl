// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package query

import (
	"enhanced_riskctrl/api/query"
	"enhanced_riskctrl/internal/conf"
	"math/big"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)
var bigZero = big.NewInt(0)
type ControllerV1 struct{
	redis  *gredis.Redis
	//
	retentionDataDur time.Duration
	enhanced_riskctrl *mpcdao.EnhancedRiskCtrl
	///
}

func NewV1() query.IQueryV1 {
	///
	t, err := gtime.ParseDuration(conf.Config.Cache.RetentionDataTime)
	if err != nil {
		panic(err)
	}
	
	///
	r := g.Redis("aggRiskCtrl")
	_, err = r.Conn(gctx.GetInitCtx())
	if err != nil {
		panic(err)
	}
	return &ControllerV1{
		retentionDataDur: t,
		redis:  r,
		enhanced_riskctrl: mpcdao.NewEnhancedRiskCtrl(r,conf.Config.Cache.Duration),
	}
}

