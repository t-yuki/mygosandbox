#!/bin/bash

go version
go test -v github.com/t-yuki/mygosandbox/gotestinittwice/a
gocov test -v github.com/t-yuki/mygosandbox/gotestinittwice/a

# Output:
# go version go1.2 linux/amd64
# 2014/01/09 15:24:51 init a
# 2014/01/09 15:24:51 init b
# 2014/01/09 15:24:51 init a
# 2014/01/09 15:24:51 init a_test
# === RUN TestFunc1
# --- PASS: TestFunc1 (0.00 seconds)
# PASS
# ok      github.com/t-yuki/mygosandbox/gotestinittwice/a 0.001s


go test -v github.com/t-yuki/mygosandbox/gotestinittwice/a/c
gocov test -v github.com/t-yuki/mygosandbox/gotestinittwice/a/c

# Output:
# 2014/01/09 15:25:35 init a
# 2014/01/09 15:25:35 init b
# 2014/01/09 15:25:35 init c
# 2014/01/09 15:25:35 init c_test
# === RUN TestFunc1
# --- PASS: TestFunc1 (0.00 seconds)
# PASS
# ok      github.com/t-yuki/mygosandbox/gotestinittwice/a/c       0.001s
