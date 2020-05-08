package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	laddr := net.TCPAddr{
		Port: 8080,
	}

	lis, err := net.ListenTCP("tcp", &laddr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.AcceptTCP()
		if err != nil {
			panic(err)
		}

		go func(conn *net.TCPConn) {
			for {
				buf := make([]byte, 1024)
				rc, err := conn.Read(buf)
				if err != nil {
					if err == io.EOF {
						log.Println("conn close")
						conn.Write([]byte("bye"))
						return
					}

					panic(err)
				}

				log.Printf("read msg[=%s], len[=%d]", buf, rc)

				wbuf := []byte(fmt.Sprintf("hello, %s", buf[:rc]))

				wc, err := conn.Write(wbuf)
				if err != nil {
					panic(err)
				}

				log.Printf("write msg[=%s], len[=%d]", wbuf, wc)
			}
		}(conn)
	}
}
