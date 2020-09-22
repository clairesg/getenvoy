package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

var (
	connectionCounterName = "proxy_wasm_go.connection_counter"
	counter               proxywasm.MetricCounter
)

type context struct{ proxywasm.DefaultContext }

func main() {
	proxywasm.SetNewRootContext(func(contextID uint32) proxywasm.RootContext { return context{} })
	proxywasm.SetNewStreamContext(func(contextID uint32) proxywasm.StreamContext { return context{} })
}

func (ctx context) OnVMStart(int) bool {
	var err error
	counter, err = proxywasm.DefineCounterMetric(connectionCounterName)
	if err != nil {
		proxywasm.LogCritical("failed to initialize connection counter: ", err.Error())
	}
	return true
}

func (ctx context) OnNewConnection() types.Action {
	proxywasm.LogInfo("new connection!")
	return types.ActionContinue
}

func (ctx context) OnDone() bool {
	err := counter.Increment(1)
	if err != nil {
		proxywasm.LogCritical("failed to increment connection counter: ", err.Error())
	}
	proxywasm.LogInfo("connection complete!")
	return true
}
