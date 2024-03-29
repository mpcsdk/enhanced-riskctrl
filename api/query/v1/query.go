package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type QueryCntReq struct {
	g.Meta   `path:"/queryCnt" tags:"query" method:"post" summary:"You first hello api"`
	ChainId  int64  `json:"chainId"`
	From     string `json:"from"`
	To       string `json:"to"`
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
type QuerySumReq struct {
	g.Meta   `path:"/querySum" tags:"querySum" method:"post" summary:"You first hello api"`
	ChainId  int64  `json:"chainId"`
	From     string `json:"from"`
	To       string `json:"to"`
	Contract string `json:"contract"`
	///
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
	//
}
type QuerySumRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Result interface{} `json:"result"`
}

///

type QueryReq struct {
	g.Meta   `path:"/query" tags:"query" method:"post" summary:"You first hello api"`
	From     string `json:"from"`
	To       string `json:"to"`
	Contract string `json:"contract"`
	///
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
	//
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
type QueryRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Result interface{} `json:"result"`
}
