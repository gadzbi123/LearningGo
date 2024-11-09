package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to dial to server: %v\n", err)
	}
	var buff = make([]byte, 512)
	for {
		fmt.Printf("Write something to server: ")
		n, err := os.Stdin.Read(buff)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
			} else {
				log.Fatalf("Failed to read to buffer: %v\n", err)
			}
		}
		var data = buff[:n]
		_, err = conn.Write(data)
		if err != nil {
			log.Fatalf("Could not write data: %v\n", err)
		}
		n, err = conn.Read(buff)
		if err != nil {
			log.Fatalf("Failed to read from conn: %v\n", err)
		}
		data = buff[:n]
		fmt.Printf("Recieved from server: %v\n", string(data))
	}
}
