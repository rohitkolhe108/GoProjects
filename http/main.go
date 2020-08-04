package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	//fmt.Println(resp)
	//read byte type, this uses Reader interface
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//use write byte, this uses Writer interface
	//io.Copy(os.Stdout, resp.Body)

	//custom writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

//custom writer
func (logWriter) Write(bs []byte) (int, error) {
	//return 1, nil
	fmt.Println(string(bs))
	fmt.Println("bytes length:", len(bs))
	return len(bs), nil

}
