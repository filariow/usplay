#!/bin/bash

docker-compose -f docker-compose.yaml -p ci up -d --build
docker logs -f ci_test_1
docker wait ci_test_1
res=$?
docker-compose -f docker-compose.yaml -p ci down

exit $res