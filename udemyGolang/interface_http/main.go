package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// custom writer
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	lw := logWriter{}
	//io.Copy(os.Stdout, res.Body)
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
