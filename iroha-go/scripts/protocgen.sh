#!/usr/bin/env bash

set -eo pipefail

proto_dirs=$(find ./schema -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  protoc \
  -I "./schema" \
  -I "./third_party/proto" \
  --go_out=./iroha.generated/protocol \
  --go_opt=paths=source_relative \
  --go-grpc_out=./iroha.generated/protocol \
  --go-grpc_opt=require_unimplemented_servers=false,paths=source_relative \
  $(find "${dir}" -maxdepth 1 -name '*.proto')
done
