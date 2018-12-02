#!/usr/bin/env bash
readonly NAME=cipher
readonly ROOT=$(dirname $0)/$NAME

$ROOT/bin/cipher $@
