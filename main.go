package main

import (
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "enhanced-riskctrl/internal/packed"

	_ "enhanced-riskctrl/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"enhanced-riskctrl/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
