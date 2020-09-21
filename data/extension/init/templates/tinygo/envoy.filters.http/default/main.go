package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {
	proxywasm.SetNewHttpContext(newContext)
}

type httpFilterContext struct {
	id uint32
	proxywasm.DefaultContext
}

func newContext(contextID uint32) proxywasm.HttpContext {
	return &httpFilterContext{id: contextID}
}

func (ctx *httpFilterContext) OnHttpRequestHeaders(int, bool) types.Action {
	// TODO:
	return types.ActionContinue
}

func (ctx *httpFilterContext) OnHttpResponseHeaders(int, bool) types.Action {
	// TODO:
	return types.ActionContinue
}
