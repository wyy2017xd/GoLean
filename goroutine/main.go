package main

import (
	"log"
	"time"
	"math/rand"
)

func printStr(str string, times int) {
	for i := 0; i < times; i++ {
		log.Println(str)
		d := time.Second * time.Duration(rand.Intn(5))/2
		time.Sleep(d)//sleep d second
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) //set rand seed
	log.SetFlags(0) // cancel log format print
	go printStr("hello", 10)
	go printStr("word", 10)
	time.Sleep(10 * time.Second) //sleep 2 s

}
