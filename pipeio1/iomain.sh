#!/bin/bash

bash -c "go run iomain.go $*; sleep 60; echo x" >> test.log 2>>err.log
