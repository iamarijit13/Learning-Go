package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// n, err := resp.Body.Read(bs)

	// if err != nil {
	// 	fmt.Println(n)
	// 	fmt.Println(string(bs))
	// }
	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many byte size: ", len(bs))
	return len(bs), nil
}
