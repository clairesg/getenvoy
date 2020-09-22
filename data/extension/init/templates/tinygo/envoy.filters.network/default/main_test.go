package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxytest"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func newStreamContext(uint32) proxywasm.StreamContext {
	return context{}
}

func TestNetwork_OnNewConnection(t *testing.T) {
	host, done := proxytest.NewNetworkFilterHost(newStreamContext)
	defer done() // release the host emulation lock so that other test cases can insert their own host emulation

	_ = host.InitConnection() // OnNewConnection is called

	logs := host.GetLogs(types.LogLevelInfo) // retrieve logs emitted to Envoy
	assert.Equal(t, logs[0], "new connection!")
}

func TestNetwork_counter(t *testing.T) {
	host, done := proxytest.NewNetworkFilterHost(newStreamContext)
	defer done() // release the host emulation lock so that other test cases can insert their own host emulation

	context{}.OnVMStart(0) // init metric

	contextID := host.InitConnection()
	host.CompleteConnection(contextID) // call OnDone on contextID -> increment the connection counter

	logs := host.GetLogs(types.LogLevelInfo)
	require.Greater(t, len(logs), 0)

	assert.Equal(t, "connection complete!", logs[len(logs)-1])
	actual, err := counter.Get()
	require.NoError(t, err)
	assert.Equal(t, uint64(1), actual)
}
