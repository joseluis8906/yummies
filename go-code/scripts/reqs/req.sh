#!/bin/bash
#
PROTO=$(cat $1 | yq '.proto')
METHOD=$(cat $1 | yq '.method')
METADATA=$(cat $1 | yq '.metadata')
REQUEST=$(cat $1 | yq '.request' -o=json)

grpcurl -import-path ../../../protobuf -proto "$PROTO" -H "$METADATA" -d "$REQUEST" -plaintext localhost:50051 $METHOD | jq 
