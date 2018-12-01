#!/usr/bin/env bash
readonly NAME=packer
readonly ROOT=$(dirname $0)/$NAME

python3 $ROOT/packer.py $@
