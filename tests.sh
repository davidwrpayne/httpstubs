#!/bin/bash
set -e
go install github.com/brentdrich/httpstubs
httpstubs &
PID=$!
trap "kill $PID" INT TERM
sleep 1 # some time to boot

echo "Test 1 - Included in cassette:"
curl -XGET -v 0.0.0.0:3000/hello
echo
echo "----------"
echo "Test 2 - Not Included in cassette:"
curl -XGET -v 0.0.0.0:3000/goodbye
echo