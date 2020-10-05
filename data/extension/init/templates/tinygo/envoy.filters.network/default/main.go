package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

var (
	connectionCounterName = "my_network_filter.connection_counter"
	counter               proxywasm.MetricCounter
)

func main() {
	proxywasm.SetNewRootContext(newRootContext)
	proxywasm.SetNewStreamContext(newStreamContext)
}

type rootContext struct {
	proxywasm.DefaultRootContext
	contextID uint32
}

func newRootContext(rootContextID uint32) proxywasm.RootContext {
	return &rootContext{contextID: rootContextID}
}

func (ctx *rootContext) OnVMStart(vmConfigurationSize int) bool {
	counter = proxywasm.DefineCounterMetric(connectionCounterName)
	return true
}

type streamContext struct {
	// we must embed the default context so that you need not to reimplement all the methods by yourself
	proxywasm.DefaultStreamContext
	rootContextID, contextID uint32
}

func newStreamContext(rootContextID, contextID uint32) proxywasm.StreamContext {
	return &streamContext{contextID: contextID, rootContextID: rootContextID}
}

func (ctx *streamContext) OnNewConnection() types.Action {
	proxywasm.LogInfo("new connection!")
	return types.ActionContinue
}

func (ctx *streamContext) OnStreamDone() {
	counter.Increment(1)
	proxywasm.LogInfo("connection complete!")
}
