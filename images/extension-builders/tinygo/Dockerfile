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

# TODO: use 1.6.0 release package after the WASI supported is released
#   this corresponds to https://github.com/tinygo-org/tinygo/commit/f50ad3585d084b17f7754f4b3cb0d42661fee036
RUN wget https://19227-136505169-gh.circle-artifacts.com/0/tmp/tinygo_amd64.deb
RUN dpkg -i tinygo_amd64.deb && rm tinygo_amd64.deb

ENV GOCACHE=/source/.gocache
ENV GOMODCACHE=/source/.gomodcache
ENV XDG_CACHE_HOME=/source/.cache

COPY ./entrypoint.sh /usr/local/getenvoy/extension/builder/entrypoint.sh
COPY ./tinygo/commands.sh /usr/local/getenvoy/extension/builder/commands.sh
ENTRYPOINT ["/usr/local/getenvoy/extension/builder/entrypoint.sh"]
CMD ["--help"]
