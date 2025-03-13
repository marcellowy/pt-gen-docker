package hello

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"os/exec"

	"pt-gen-docker/api/hello/v1"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {

	if req.Sid == "" {
		return
	}

	cmd := exec.Command("node", "pt-gen-cfworker-master/main.js", req.Sid)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	g.RequestFromCtx(ctx).Response.WriteJsonExit(output)
	//fmt.Println(string(output))
	//g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
