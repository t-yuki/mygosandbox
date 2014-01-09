package b

import (
	"log"

	_ "github.com/t-yuki/mygosandbox/gotestinittwice/a"
)

func init() {
	log.Println("init b")
}
