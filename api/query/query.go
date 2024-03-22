// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package query

import (
	"context"
	
	"enhanced-riskctrl/api/query/v1"
)

type IQueryV1 interface {
	QueryCnt(ctx context.Context, req *v1.QueryCntReq) (res *v1.QueryCntRes, err error)
}


