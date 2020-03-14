package main

import (
	"GoLean/myHashMap"
	"fmt"
)
func main() {
	myMap := myHashMap.CreateHashMap()
	myMap.AddKeyValue("1", "wyy")
	myMap.AddKeyValue("2", "wjl")
	myMap.AddKeyValue("3", "wjp")
	fmt.Println(myMap)
}
