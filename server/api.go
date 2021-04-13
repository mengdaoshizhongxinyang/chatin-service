package server

import (
	"Chatin/logic"
	"github.com/tidwall/gjson"
)

type resultGetter interface {
	Get(string) gjson.Result
}
type handler func(action string,p resultGetter) logic.MSG
type apiCaller struct {
	handlers []handler
}
