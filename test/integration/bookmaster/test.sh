#!/bin/bash

function run_containerizer_tests() {
    docker-compose -f docker-compose.yaml -p ci up -d --build
    docker logs -f ci_test_1
    docker wait ci_test_1
    res=$?
    docker-compose -f docker-compose.yaml -p ci down

    return $res
}

function local_cmd() {
    if [ -z $2 ]; then
        echo "up or down?"
    elif [ $2 == "up" ]; then
        docker-compose -f docker-compose.yaml -p ci up -d mongoitest
    elif [ $2 == "down" ]; then
        docker-compose -f docker-compose.yaml -p ci down
    else
        echo "unrecognized argument $1 $2"
    fi
}


if [ -z $1 ]; then
    run_containerizer_tests
elif [ $1 == "local" ]; then
    local_cmd $1 $2
else
    echo "Unrecognized argument $1"
fi