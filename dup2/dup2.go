package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) < 0 {
		return
	}
	for _, arg := range files {
		f, err := os.Open(arg)
		if err == nil {
			countLines(counts, f)
		}

	}
	for k, v := range counts {
		if v > 0 {
			fmt.Println(k)
		}
	}
}

func countLines(counts map[string]int, f *os.File) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}