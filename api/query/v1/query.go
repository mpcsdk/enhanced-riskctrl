package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type QueryCntReq struct {
	g.Meta   `path:"/query" tags:"query" method:"post" summary:"You first hello api"`
	ChainId  int64  `json:"chainId"`
	FromAddr string `json:"fromAddr"`
	ToAddr   string `json:"toAddr"`
	Contract string `json:"contract"`
	///
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
	//
}
type QueryCntRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Result interface{} `json:"result"`
}

// //
