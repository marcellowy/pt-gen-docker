package cmd

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"pt-gen-docker/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
				)
			})

			cmd := exec.Command("node", "pt-gen-cfworker-master/main.js", "34429795")
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println(string(output))

			s.Run()
			return nil
		},
	}
)
