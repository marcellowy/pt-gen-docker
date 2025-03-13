package main

import (
	_ "pt-gen-docker/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"pt-gen-docker/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
