package c

import (
	"log"

	_ "github.com/t-yuki/mygosandbox/gotestinittwice/a"
	_ "github.com/t-yuki/mygosandbox/gotestinittwice/a/b"
)

func init() {
	log.Println("init c")
}
