#!/bin/bash

go test || exit 1
for i in `seq 8`; do
       	GOMAXPROCS=`echo 2^$i|bc -l` go test || exit 1
done
