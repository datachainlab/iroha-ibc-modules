#!/usr/bin/env bash

set -eo pipefail

proto_dirs=$(find ./schema -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  protoc \
  -I "./schema" \
  -I "./third_party/proto" \
  --go_out=./iroha.generated/protocol \
  --go_opt=module=iroha.generated/protocol \
  --go-grpc_out=. \
  $(find "${dir}" -maxdepth 1 -name '*.proto')
done
