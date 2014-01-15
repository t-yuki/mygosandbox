#!/bin/bash

#

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


# go test -c github.com/t-yuki/mygosandbox/gotestinittwice/a
# objdump -t -T a.test | grep "a.init"
# Output:
# objdump: a.test: not a dynamic object
# 00000000006538a5 g     O .noptrbss      0000000000000001 github.com/t-yuki/mygosandbox/gotestinittwice/a.initdone・
# 00000000006538a0 g     O .noptrbss      0000000000000001 _/home/t-yuki/gopath.t-yuki/src/github.com/t-yuki/mygosandbox/gotestinittwice/a.initdone・
# 000000000041c230 g     F .text  00000000000002ba runtime.symtabinit
# 0000000000493680 g     F .text  00000000000000b5 _/home/t-yuki/gopath.t-yuki/src/github.com/t-yuki/mygosandbox/gotestinittwice/a.init・1
# 0000000000493740 g     F .text  000000000000004b _/home/t-yuki/gopath.t-yuki/src/github.com/t-yuki/mygosandbox/gotestinittwice/a.init
# 00000000004af820 g     F .text  00000000000000b5 github.com/t-yuki/mygosandbox/gotestinittwice/a.init・1
# 00000000004af8e0 g     F .text  000000000000004b github.com/t-yuki/mygosandbox/gotestinittwice/a.init
