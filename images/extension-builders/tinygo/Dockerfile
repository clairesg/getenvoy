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

#
# Builder image for Envoy Wasm extensions written in Go.
#

FROM golang:1.15

ENV TINYGO_VERSION=0.15.0

RUN wget https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo_${TINYGO_VERSION}_amd64.deb
RUN dpkg -i tinygo_${TINYGO_VERSION}_amd64.deb && rm tinygo_${TINYGO_VERSION}_amd64.deb

ENV GOCACHE=/source/.gocache
ENV GOMODCACHE=/source/.gomodcache
ENV XDG_CACHE_HOME=/source/.cache

ENV TINYGO_SDK_NAME=github.com/tetratelabs/proxy-wasm-go-sdk
ENV TINYGO_SDK_VERSION=v0.0.2
ENV TINYGO_SDK_PATH=${GOMODCACHE}/${TINYGO_SDK_NAME}@${TINYGO_SDK_VERSION}

RUN mkdir -p ${TINYGO_SDK_PATH} && git clone https://${TINYGO_SDK_NAME} -b ${TINYGO_SDK_VERSION} ${TINYGO_SDK_PATH}

COPY ./entrypoint.sh /usr/local/getenvoy/extension/builder/entrypoint.sh
COPY ./tinygo/commands.sh /usr/local/getenvoy/extension/builder/commands.sh
ENTRYPOINT ["/usr/local/getenvoy/extension/builder/entrypoint.sh"]
CMD ["--help"]
