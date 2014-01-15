package socket_test

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"testing"
	"time"
)

type Packet struct {
	A int64
	B string
}

func serverProc(l net.Listener) {
	for {
		fd, err := l.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			return
		}

		log.Printf("accept: %v", fd)
		echoServer(fd)
	}
}

func echoServer(c net.Conn) {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan *Packet, 100)
	go func() {
		dec := gob.NewDecoder(c)
		for {
			p := &Packet{}
			if err := dec.Decode(p); err != nil {
				log.Printf("server receive error: %v", err)
				break
			}
			log.Printf("Server received: %v", p)
			time.Sleep(1e5)
			ch <- p
		}
		close(ch)
		wg.Done()
	}()

	go func() {
		enc := gob.NewEncoder(c)
		for {
			p, open := <-ch
			if p == nil && !open {
				log.Printf("server close")
				break
			}
			if err := enc.Encode(p); err != nil {
				log.Printf("server send error: %v", err)
				break
			}
			log.Printf("Server sent: %v", p)
		}
		wg.Done()
	}()
	wg.Wait()
}

func Test(t *testing.T) {
	os.Remove("/tmp/echo.sock")
	l, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		log.Printf("listen error: %v", err)
		return
	}

	go serverProc(l)

	var wg sync.WaitGroup
	wg.Add(3)
	go clientProc(1, &wg)
	go clientProc(2, &wg)
	go clientProc(3, &wg)
	wg.Wait()
}

func clientProc(id int, wait *sync.WaitGroup) {
	defer wait.Done()
	for i := 0; i < 3; i++ {
		c, err := net.Dial("unix", "/tmp/echo.sock")
		if err != nil {
			panic(err)
		}
		log.Printf("begin client %d - %d", id, i)
		dec := gob.NewDecoder(c)
		enc := gob.NewEncoder(c)

		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			for j := 0; j < 4; j++ {
				p := &Packet{}
				p.A = int64(j)
				p.B = fmt.Sprintf("p:%d - %d - %d", id, i, j)
				if err := enc.Encode(p); err != nil {
					log.Print(err)
					break
				}
				log.Printf("Sent: %v", p)
				time.Sleep(1e6)
			}
			log.Printf("close client %d - %d", id, i)
			c.Close()
			wg.Done()
		}()
		go func() {
			for j := 0; j < 10; j++ {
				p := &Packet{}
				if err := dec.Decode(p); err != nil {
					log.Print(err)
					break
				}
				log.Printf("Received: %v", p)
				time.Sleep(1e6)
			}
			wg.Done()
		}()

		//time.Sleep(1e7)
		//log.Printf("close client %d - %d", id, i)
		//c.Close()
		wg.Wait()
		log.Printf("end client %d - %d", id, i)
	}
}
