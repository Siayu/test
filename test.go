package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	userFile := "hello.txt"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	//for {
	n, _ := fl.Read(buf)
	if 0 == n {
		fmt.Println("end")
		//break
	}

	s := string(buf)
	fmt.Println(strings.Split(s, "\n"))
	m := strings.Split(s, "\n")

	//for k, v:= range buf {
	//	fmt.Printf("k: %d, v: %s", k, v)
	//}
	fmt.Println(m)
	x:=0
	for x<len(m){
		if strings.ContainsAny(m[x],"a"){
		fmt.Println(m[x])
		}
		x++
	}


	//os.Stdout.Write(buf[:n])
	//}

}