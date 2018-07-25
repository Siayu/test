package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
    //新建一个文档 newcomment
	userFile := "newcomment.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()

	//读取 comment文档
	userFile1:= "comment.txt"
	fl, err := os.Open(userFile1)
	if err != nil {
		fmt.Println(userFile1, err)
		return
	}
	defer fl.Close()
	buf := make([]byte,1321312)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
	}
	s := string(buf)
	fmt.Println(strings.Split(s, "\n"))
	m := strings.Split(s, "\n")
    x:=0
    w:=0
	for x<len(m){
		if strings.ContainsAny(m[x],"客服"){
			fmt.Println(m[x])
			r := strconv.Itoa(x+1)
			fout.WriteString(r+","+m[x]+"\n")
			w++
		}
		x++
	}
	fmt.Println(w)
}
