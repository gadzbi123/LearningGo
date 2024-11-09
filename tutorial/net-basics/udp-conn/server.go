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

	conn, err := net.ListenPacket("udp", port)
	if err != nil {
		log.Fatalf("Failed to listen on connection: %v\n", err)
	}
	var buff = make([]byte, 512)
	for {
		n, addr, err := conn.ReadFrom(buff)
		log.Printf("packet recieved from: %v\n", addr)
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

		_, err = conn.WriteTo(bytes.ToUpper(data), addr)
		if err != nil {
			log.Fatalf("Failed to write to conn: %v\n", err)
		}
	}
}
