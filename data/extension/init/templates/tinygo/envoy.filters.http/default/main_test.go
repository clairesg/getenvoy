package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxytest"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func TestHttpHeaders_OnHttpRequestHeaders(t *testing.T) {
	host, done := proxytest.NewHttpFilterHost(func(contextID uint32) proxywasm.HttpContext {
		return &context{}
	})
	defer done()
	id := host.InitContext()

	hs := [][2]string{
		{"key1", "value1"},
		{"key2", "value2"},
	}
	host.PutRequestHeaders(id, hs) // call OnHttpRequestHeaders

	logs := host.GetLogs(types.LogLevelInfo)
	require.Greater(t, len(logs), 2)

	assert.Equal(t, "key2: value2", logs[len(logs)-1])
	assert.Equal(t, "key1: value1", logs[len(logs)-2])
	assert.Equal(t, "observing request headers", logs[len(logs)-3])
}
