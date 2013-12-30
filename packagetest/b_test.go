package a

import (
	package1 "./package1.packagetest"
	package2 "./package2.packagetest"
	"testing"
)

func TestX(t *testing.T) {
	package1.Pub()
	package2.Pub()
}
