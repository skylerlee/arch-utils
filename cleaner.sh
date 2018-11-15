#!/usr/bin/env bash
readonly NAME=cleaner
readonly ROOT=$(dirname $0)/$NAME

pushd $ROOT > /dev/null

./bin/cleanfs

popd > /dev/null
