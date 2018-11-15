#!/usr/bin/env bash
readonly NAME=proxy
readonly ROOT=$(dirname $0)/$NAME

pushd $ROOT > /dev/null

nohup ss-local -c conf/shadowsocks.json > log/error.log 2>&1 &

popd > /dev/null
