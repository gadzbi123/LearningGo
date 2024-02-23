package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(arr []byte) (int, error) {
	n, err := (reader.r.Read(arr))

	if err == io.EOF {
		return 0, io.EOF
	}

	for i, _ := range arr {
		if arr[i] >= 'a' && arr[i] <= 'm' || arr[i] >= 'A' && arr[i] <= 'M' {
			arr[i] += 13
		} else if arr[i] >= 'n' && arr[i] <= 'z' || arr[i] >= 'N' && arr[i] <= 'Z' {
			arr[i] -= 13
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	for {
		arr := make([]byte, 8)
		_, err := r.Read(arr)
		if err != nil {
			break
		}
		fmt.Println(string(arr))
	}
	io.Copy(os.Stdout, s)
}
