package main

import (
	"os"
	"io/ioutil"
	"strings"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for k, v := range counts {
		if  v > 0 {
			fmt.Println(k)
		}
	}
}
