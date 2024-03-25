package main

import (
	"enhanced-riskctrl/internal/conf"
	_ "enhanced-riskctrl/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2/frame/g"

	_ "enhanced-riskctrl/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"enhanced-riskctrl/internal/cmd"
)

func main() {
	g.Log().SetAsync(true)
	g.Log().SetWriterColorEnable(true)

	ctx := gctx.New()

	if conf.Config.Jaeger.Enable {
		tp, err := jaeger.Init(conf.Config.Server.Name, conf.Config.Jaeger.Url)
		if err != nil {
			g.Log().Error(ctx, err)
		}
		defer tp.Shutdown(ctx)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
