#!/bin/bash
#
etcdctl --endpoints=yummies.local:2379 get /config/yummies.yml --print-value-only | yq
