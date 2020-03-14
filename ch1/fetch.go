package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch is %v\n", err)
			os.Exit(0)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprint(os.Stderr, "fecth reading %v\n", err)
		}
		fmt.Println(resp.Header, b)
		resp.Body.Close()
	}
}
