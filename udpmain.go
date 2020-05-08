package main

import (
	"log"
	"net"
)

func main() {

	laddr := net.UDPAddr{
		Port: 18080,
	}

	conn, err := net.ListenUDP("udp", &laddr)
	if err != nil {
		panic(err)
	}

	for {
		readBuf := make([]byte, 1024)
		rc, clientAddr, err := conn.ReadFromUDP(readBuf)
		if err != nil {
			panic(err)
		}

		log.Printf("read from client addr[=%s], msg[=%s], len[=%d]", clientAddr, readBuf, rc)

		writeBuf := append([]byte("hello: "), readBuf[:rc]...)
		wc, err := conn.WriteToUDP(writeBuf, clientAddr)
		if err != nil {
			panic(err)
		}

		log.Printf("write to msg[=%s], len[=%d]", writeBuf, wc)
	}
}
