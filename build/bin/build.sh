#!/bin/bash

# check inputs
if [ $# -eq 0 ]
  then
    echo "no arguments passed to buildall.sh script"
    exit 1
fi

# build all given projects
for i in $@
do
    make bin PRJ_TARGET=$i TARGET=all
done

# cleanup
make clean