package main

import (
	"fmt"
	"io"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func GenerateData(w io.Writer) {
	data := []byte("lubie jechac autem")
	size := 4
	for i := 0; i < len(data); i += size {
		end := i + size
		if end > len(data) {
			end = len(data)
		}
		n, err := w.Write(data[i:end])
		if err != nil {
			Printfln("Error: %v", err)
		}
		Printfln("Write Chunk: \"%v\", bytes received: %v", string(data[i:i+n]), n)
	}
	if closer, ok := w.(io.Closer); ok {
		closer.Close()
	}
}
func ReadData(r io.Reader) {
	chunk := make([]byte, 2)
	for {
		n, err := r.Read(chunk)
		if err == io.EOF {
			break
		}
		Printfln("Read chunk: %v, len: %v, error: %v", string(chunk[:n]), n, err)
	}
}
func main() {
	pipeReader, pipeWriter := io.Pipe()
	go func() {
		GenerateData(pipeWriter)
	}()
	ReadData(pipeReader)
}
