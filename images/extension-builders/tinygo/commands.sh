#!/usr/bin/env bash

# Copyright 2020 Tetrate
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#########################################################################
# Build Wasm extension and copy *.wasm file to a given location.
# Globals:
#   CARGO_TARGET_DIR
#   GETENVOY_WORKSPACE_DIR
# Arguments:
#   Path relative to the workspace root to copy *.wasm file to.
#########################################################################
extension_build()  {
	tinygo build -o "$1" -wasm-abi=generic -target wasm main.go
}

extension_test()  {
	go test -tags=proxytest -v ./...
}

extension_clean()  {
	rm main.wasm || true
	rm -rf "${GOCACHE}" "${XDG_CACHE_HOME}" "${GOMODCACHE}" || true
}
