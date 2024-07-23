// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package query

import (
	"context"

	"enhanced_riskctrl/api/query/v1"
)

type IQueryV1 interface {
	QueryCnt(ctx context.Context, req *v1.QueryCntReq) (res *v1.QueryCntRes, err error)
	QuerySum(ctx context.Context, req *v1.QuerySumReq) (res *v1.QuerySumRes, err error)
	Query(ctx context.Context, req *v1.QueryReq) (res *v1.QueryRes, err error)
	State(ctx context.Context, req *v1.StateReq) (res *v1.StateRes, err error)
}
