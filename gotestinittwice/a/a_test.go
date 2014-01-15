package a_test

import (
	"log"
	"testing"

	_ "."
	_ "github.com/t-yuki/mygosandbox/gotestinittwice/a/b"
)

func init() {
	log.Println("init a_test")
}

func TestFunc1(t *testing.T) {
}
