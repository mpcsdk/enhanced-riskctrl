package query

import (
	"context"

	"github.com/mpcsdk/mpcCommon/mpcdao"

	v1 "enhanced-riskctrl/api/query/v1"
)

func (c *ControllerV1) QueryCnt(ctx context.Context, req *v1.QueryCntReq) (res *v1.QueryCntRes, err error) {
	///
	data, err := mpcdao.GetAggNft(ctx, &mpcdao.QueryAggNft{})
	if err != nil {
		return nil, err
	}
	///
	return &v1.QueryCntRes{
		Result: data.Value,
	}, nil
}
