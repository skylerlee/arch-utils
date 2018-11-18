#!/usr/bin/env bash
readonly NAME=server
readonly ROOT=$(dirname $0)/$NAME

pushd $ROOT > /dev/null

./bin/server $@

popd > /dev/null
