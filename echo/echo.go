package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//for i := 1; i < len(os.Args); i++ {
	//	fmt.Println(os.Args[i])
	//}
	fmt.Println(strings.Join(os.Args[1:], " "))

}
