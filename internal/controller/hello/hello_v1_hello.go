package hello

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"os"
	"os/exec"
	"sync"

	"pt-gen-docker/api/hello/v1"
)

var lock = sync.Mutex{}

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {

	if req.Sid == "" {
		return
	}

	lock.Lock()
	defer lock.Unlock()

	var cacheDir = "/app/cache"
	var output []byte

	// 判断是不是有缓存
	cacheFile := cacheDir + "/" + req.Sid + ".json"
	if _, err = os.Stat(cacheFile); err != nil {
		fmt.Println("cache file not exists")
		cmd := exec.Command("node", "pt-gen-cfworker-master/main.js", req.Sid)
		output, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
			return
		}
		if err = os.WriteFile(cacheFile, output, os.ModePerm); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("cache file exists")
		if output, err = os.ReadFile(cacheFile); err != nil {
			fmt.Println(err)
		}
	}

	g.RequestFromCtx(ctx).Response.WriteJsonExit(output)
	//fmt.Println(string(output))
	//g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
