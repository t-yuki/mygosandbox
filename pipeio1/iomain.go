package main

import (
	"flag"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	path1 := flag.String("path1", "", "")
	path2 := flag.String("path2", "", "")
	count := flag.Int("count", 0, "")
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(1)
	in, err := os.OpenFile(*path2, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	out, err := os.OpenFile(*path1, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	go func() {
		for i := 0; i < *count; i++ {
			if n, err := out.Write([]byte("abc")); err != nil {
				log.Printf("wr: %v", err)
				break
			} else if n == 0 {
				log.Printf("wr: 0")
				break
			}
		}
		out.Close()
		wg.Done()
	}()
	for {
		buf := make([]byte, 256)
		if n, err := io.ReadFull(in, buf); err != nil {
			log.Printf("rd: %v", err)
			break
		} else if n == 0 {
			log.Printf("rd: 0")
			break
		}
	}
	in.Close()
	out.Close()
	wg.Wait()
}
