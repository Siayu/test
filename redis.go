package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

type response1 struct{
	code int
	msg string
	data string
}

func main() {

	var a response1
	//resp, err := http.Get("http://localhost:6379/submituser")
	//defer resp.Body.Close()
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	exists ,err := redis.Bool(c.Do("EXISTS","ann"))
	if err != nil{
		fmt.Println(err)
	}
	if exists {
		a.code = 0
		a.msg = "ok"
	}else{
		a.code= 1001
		a.msg= "错误的原因"
	}


	var b response1
	//resp, err := http.Get("http://localhost:6379/userquery")
	//defer resp.Body.Close()
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	exists ,err := redis.Bool(c.Do("EXISTS","uid"))
	if err != nil{
		fmt.Println(err)
	}

	if exists {
		a.code = 0
		a.msg = "ok"
	}else{
		a.code= 1001
		a.msg= "错误的原因"
	}
}
