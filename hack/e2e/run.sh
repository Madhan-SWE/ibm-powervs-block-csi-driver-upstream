#!/bin/bash

# Copyright 2022 The Kubernetes Authors.
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

set -euo pipefail

BASE_DIR=$(dirname "$(realpath "${BASH_SOURCE[0]}")")
source "${BASE_DIR}"/util.sh

TEST_DIR=${BASE_DIR}/csi-test-artifacts
BIN_DIR=${TEST_DIR}/bin

TEST_PATH=${TEST_PATH:-"./tests/e2e/..."}
ARTIFACTS=${ARTIFACTS:-"${TEST_DIR}/artifacts"}
GINKGO_FOCUS=${GINKGO_FOCUS:-"\[powervs-csi-e2e\]"}
GINKGO_SKIP=${GINKGO_SKIP:-"\[Disruptive\]"}
GINKGO_NODES=${GINKGO_NODES:-1}
TEST_EXTRA_FLAGS=${TEST_EXTRA_FLAGS:-}


loudecho "Installing ginkgo to ${BIN_DIR}"

GINKGO_BIN=${BIN_DIR}/ginkgo
if [[ ! -e ${GINKGO_BIN} ]]; then
  pushd /tmp
  CGO_ENABLED=0 GOPATH=${TEST_DIR} GOBIN=${BIN_DIR} GO111MODULE=on go install github.com/onsi/ginkgo/ginkgo@v1.16.5
  popd
fi

loudecho "Testing focus ${GINKGO_FOCUS}"
eval "EXPANDED_TEST_EXTRA_FLAGS=$TEST_EXTRA_FLAGS"
set -x
set +e
CGO_ENABLED=0 ${GINKGO_BIN} -p -nodes="${GINKGO_NODES}" -v --focus="${GINKGO_FOCUS}" --skip="${GINKGO_SKIP}" "${TEST_PATH}" -report-dir="${ARTIFACTS}" "${EXPANDED_TEST_EXTRA_FLAGS}"
TEST_PASSED=$?
set -e
set +x
loudecho "TEST_PASSED: ${TEST_PASSED}"

if [ "${TEST_PASSED}" == 0 ]; then
    loudecho "Tests passed"
else
    loudecho "One of the tests failed"
    exit 1
fi
