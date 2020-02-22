package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	num = 10
	var  str1 string
	var str2 string
	str1 = fmt.Sprintf("%d", num)
	fmt.Println(str1)
	str2 = strconv.FormatInt((int64(num)), 10)
	fmt.Println(str2)
	str3 := strconv.Itoa(num)
	fmt.Println(str3)
	var str4 string
	str4 = "20"
	num1, _ := strconv.ParseInt(str4, 10, 0)//0 int
	fmt.Println(num1)
}
