#!/bin/bash
#
if [ "$#" -eq "0" ]; then
    echo "usage:"
    echo "./import.sh <file>.yml"
    exit -1
fi

filename=$(basename -- "$1")
collection="${filename%.*}"


yq $1 -o json > /tmp/data.json
isArray=$(cat /tmp/data.json | jq 'if type=="array" then "yes" else "no" end' | tr -d '"')

if [[ $isArray == "yes" ]]; then
    mongoimport "mongodb://yummies:yummies@yummies.local:27017/yummies?authMechanism=PLAIN" --collection="$collection" --file=/tmp/data.json --jsonArray
else
    mongoimport "mongodb://yummies:yummies@yummies.local:27017/yummies?authMechanism=PLAIN" --collection="$collection" --file=/tmp/data.json
fi

