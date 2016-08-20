#!/bin/bash
go install github.com/brentdrich/httpstubs
httpstubs &
PID=$!
trap "kill -9 $PID" INT TERM
curl -XGET -v 0.0.0.0:3000/
echo