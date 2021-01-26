package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	log.Println("start tcp listen")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("listener.Accept(\"%s\") error(%v)", ln.Addr().String(), err)
			return
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()

	readCh := make(chan []byte, 10)
	defer close(readCh)
	sendCh := make(chan []byte, 10)
	defer close(sendCh)

	go func(readCh <-chan []byte, sendCh chan<- []byte) {
		for msg := range readCh {
			// TODO 解析数据包
			log.Printf("recv client msg: %v", string(msg))
			sendCh <- msg
		}
	}(readCh, sendCh)

	go func(conn net.Conn, sendCh <-chan []byte) {
		for msg := range sendCh {
			conn.Write(msg)
		}
	}(conn, sendCh)

	reader := bufio.NewReader(conn)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Printf("read error: %v", err)
			}
			return
		}
		readCh <- line
	}
}
