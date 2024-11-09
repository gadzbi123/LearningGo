package main

import (
	"bytes"
	"io"
	"log"
	"net"
)

func main() {
	port := ":8080"
	log.Printf("Server is running at port: %v\n", port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on connection: %v\n", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Could not accept conn: %v\n", err)
	}
	log.Printf("connection accepted from: %v\n", conn.RemoteAddr())
	var buff = make([]byte, 512)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			switch err {
			case io.EOF:
				log.Println("EOF")
			case net.ErrClosed:
				log.Println("Conn closed")
			default:
				log.Fatalf("Failed on reading from conn: %v\n", err)
			}
		}
		var data = buff[:n]

		_, err = conn.Write(bytes.ToUpper(data))
		if err != nil {
			log.Fatalf("Failed to write to conn: %v\n", err)
		}
	}
}
