package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		//fmt.Println("1",s,time.Now().Unix())
		time.Sleep(1 * time.Second)
		fmt.Println(s, time.Now().Unix())
	}
}

func main() {
	go say("world")
	say("hello")
}
