#!/bin/bash

find . -iname \*_test.go | xargs dirname | sort -u | xargs go test