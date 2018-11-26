#!/usr/bin/env bash
readonly NAME=server
readonly ROOT=$(dirname $0)/$NAME

$ROOT/bin/server $@
