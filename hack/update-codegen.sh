#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

bash "${CODEGEN_PKG}"/generate-groups.sh "deepcopy,client,informer" \
  github.com/satchel9/virtualization/pkg/client github.com/satchel9/virtualization/pkg/apis virtualization:v1 \
  --output-base "${SCRIPT_ROOT}"/../ \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt

