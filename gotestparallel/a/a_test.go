package a_test

import (
	"bytes"
	"net"
	"os/exec"
	"testing"
	"time"
)

func TestFunc1(t *testing.T) {
	_, err := net.Listen("tcp", ":1234")
	if err != nil {
		cmd := exec.Command("ps", "ux")
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
		t.Log(out.String())
		t.Fatal(err)
	}
	time.Sleep(time.Second)
}
