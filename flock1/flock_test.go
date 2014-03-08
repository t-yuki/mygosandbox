package flock1_test

import (
	"os"
	"syscall"
	"testing"
)

func TestFlock(t *testing.T) {
	fd, err := syscall.Open("flocktest.lock", syscall.O_CREAT|syscall.O_RDWR, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("flocktest.lock")
	defer syscall.Close(fd)
	if err = syscall.Flock(fd, syscall.LOCK_SH); err != nil {
		t.Fatal(err)
	}

	fd2, err := syscall.Open("flocktest.lock", syscall.O_RDWR, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer syscall.Close(fd2)

	fd3, err := syscall.Open("flocktest.lock", syscall.O_RDWR, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer syscall.Close(fd3)

	if err = syscall.Flock(fd2, syscall.LOCK_EX|syscall.LOCK_NB); err == nil {
		t.Fatal(err)
	}

	if err = syscall.Flock(fd2, syscall.LOCK_SH|syscall.LOCK_NB); err != nil {
		t.Fatal(err)
	}

	if err = syscall.Close(fd); err != nil {
		t.Fatal(err)
	}
	if err = syscall.Flock(fd3, syscall.LOCK_EX|syscall.LOCK_NB); err == nil {
		t.Fatal(err)
	}
	if err = syscall.Flock(fd2, syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		t.Fatal(err)
	}
}
