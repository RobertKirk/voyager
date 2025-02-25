#!/bin/bash

# Copyright The Voyager Authors.
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

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT=$(dirname "${BASH_SOURCE[0]}")/../../../..

source "$REPO_ROOT/hack/libbuild/common/public_image.sh"

detect_tag $REPO_ROOT/dist/.tag

IMG=haproxy
TAG=1.7.6-$TAG

build() {
    pushd $(dirname "${BASH_SOURCE}")
    wget -O kloader https://cdn.appscode.com/binaries/kloader/4.0.1/kloader-linux-amd64
    chmod +x kloader
    local cmd="docker build --pull -t appscode/$IMG:$TAG ."
    echo $cmd
    $cmd
    rm kloader
    popd
}

binary_repo $@
