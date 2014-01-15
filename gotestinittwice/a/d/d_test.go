package d_test

import (
	"log"
	"testing"

	// "." and "a/d" are the same package!
	_ "."
	_ "github.com/t-yuki/mygosandbox/gotestinittwice/a/d"
)

func init() {
	log.Println("init d_test")
}

func TestFunc1(t *testing.T) {
}
