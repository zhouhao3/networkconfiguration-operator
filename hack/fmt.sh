#!/bin/sh

set -u

for file in $(find . -path ./.git -prune -o -type f | grep -E ".*\.go$")
do
    result=$(./tools/gofmt -s -d $file)
    if [ "$result" != "" ];then
        echo "$result"
        exit 1
    fi
done




