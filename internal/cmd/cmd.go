package cmd

import (
	"context"
	"enhanced-riskctrl/internal/controller/nrpc"
	"enhanced-riskctrl/internal/controller/query"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// /
			///
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					query.NewV1(),
				)
			})
			///
			go nrpc.Run()
			s.Run()
			///

			return nil
		},
	}
)
