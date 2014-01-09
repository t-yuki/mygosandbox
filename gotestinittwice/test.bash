#!/bin/bash

go version
go test -v github.com/t-yuki/mygosandbox/gotestinittwice/...
gocov test -v github.com/t-yuki/mygosandbox/gotestinittwice/...

# the above may cause panic

# go version go1.2 linux/amd64
# 2014/01/09 15:01:53 init a
# 2014/01/09 15:01:53 init b
# 2014/01/09 15:01:53 init a
# 2014/01/09 15:01:53 init a_test
# === RUN TestFunc1
# --- PASS: TestFunc1 (0.00 seconds)
# PASS
# ok      github.com/t-yuki/mygosandbox/gotestinittwice/a 0.001s
# 2014/01/09 15:01:53 init a
# 2014/01/09 15:01:53 init b
# 2014/01/09 15:01:53 init b_test
# === RUN TestFunc1
# --- PASS: TestFunc1 (0.00 seconds)
# PASS
# ok      github.com/t-yuki/mygosandbox/gotestinittwice/a/b       0.001s
