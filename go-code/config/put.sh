#!/bin/bash
#
etcdctl --endpoints=yummies.local:2379 put /config/yummies.yml "$(cat -v ./yummies.yml)"
