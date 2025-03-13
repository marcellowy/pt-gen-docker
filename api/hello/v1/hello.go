package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/v1/ptgen" tags:"Hello" method:"post" summary:"You first hello api"`
	Sid    string `json:"sid"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
