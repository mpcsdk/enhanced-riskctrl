package receiver

import (
	"context"
	"math"
	"math/big"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

func (s *sReceiver) clean() {

}

func (s *sReceiver) aggTx(ctx context.Context, key string, tx *entity.ChainData) error {
	_, err := g.Redis().Do(ctx, "Zadd", key, tx.Ts, tx)
	if err != nil {
		return err
	}
	return nil
}

func (s *sReceiver) getTxCnt(ctx context.Context, key string, startTs, endTs int64) (int64, error) {
	v, err := g.Redis().Do(ctx, "ZCARD", key, startTs, endTs)
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

// /

func (s *sReceiver) getTxSum(ctx context.Context, key string, startTs, endTs int64) (*big.Int, error) {
	if endTs == 0 {
		endTs = math.MaxInt64
	}
	v, err := g.Redis().Do(ctx, "ZRANGE", key, startTs, endTs)
	if err != nil {
		return nil, err
	}
	//
	data := []*entity.ChainData{}
	v.Struct(&data)
	///
	sum := big.NewInt(0)

	for _, tx := range data {
		i := big.NewInt(0)
		i.SetString(tx.Value, 10)
		sum = sum.Add(sum, i)
	}
	return sum, nil
}
