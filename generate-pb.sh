#!/bin/bash
#
protos=$(fd ".proto" ./protobuf/)

for proto in $protos; do
  protoc --proto_path=protobuf --go_out=go-code/pkg/pb --go_opt=paths=source_relative --js_out=import_style=commonjs:js-code/pb --go-grpc_out=go-code/pkg/pb --go-grpc_opt=paths=source_relative --grpc-web_out=import_style=typescript,mode=grpcwebtext:js-code/pb $proto
done
