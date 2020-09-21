package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

var (
	requestCounterName = "proxy_wasm_go.request_counter"
	counter            proxywasm.MetricCounter
)

func main() {
	proxywasm.SetNewRootContext(func(uint32) proxywasm.RootContext { return &context{} })
	proxywasm.SetNewHttpContext(func(uint32) proxywasm.HttpContext { return &context{} })
}

type context struct{ proxywasm.DefaultContext }

func (ctx *context) OnVMStart(int) bool {
	var err error
	counter, err = proxywasm.DefineCounterMetric(requestCounterName)
	if err != nil {
		proxywasm.LogCritical("failed to initialize request counter: ", err.Error())
	}
	return true
}

func (ctx *context) OnHttpRequestHeaders(int, bool) types.Action {
	hs, err := proxywasm.HostCallGetHttpRequestHeaders()
	if err != nil {
		proxywasm.LogCritical("failed to get request headers: ", err.Error())
	}

	proxywasm.LogInfo("observing request headers")
	for _, h := range hs {
		proxywasm.LogInfo(h[0], ": ", h[1])
	}

	return types.ActionContinue
}

func (ctx *context) OnHttpResponseHeaders(int, bool) types.Action {
	if err := proxywasm.HostCallSetHttpResponseHeader("additional", "header"); err != nil {
		proxywasm.LogCritical(err.Error())
	}
	return types.ActionContinue
}

func (ctx *context) OnDone() bool {
	if err := counter.Increment(1); err != nil {
		proxywasm.LogCritical("failed to increment request counter: ", err.Error())
	}
	return true
}
