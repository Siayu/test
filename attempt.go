package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"encoding/json"
)

type response3 struct {
	Code int
	Msg string
	Data string
}

func main(){
	var a response3
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("error is" ,err)
	}
	var x,y,z string
	fmt.Println("Input the uid, name and age")
	fmt.Scanf("%s",&x)
	c.Do("HSET", "submituser", "uid"+x, x )
	fmt.Scanf("%s",&y)
	c.Do("HSET", "submituser", "name"+x, y)
	fmt.Scanf("%s",&z)
	c.Do("HSET", "submituser", "age"+x, z)
	c.Do("HGET", "submituser", "uid"+x)
	c.Do("HGET", "submituser", "name"+x)
	c.Do("HGET", "submituser", "age"+x)
	exists ,err := redis.Bool(c.Do("HEXIST","uid"+x))
	if err != nil{
	fmt.Println("error", err)
	}
	if exists {
		a.Code = 0
		a.Msg = "ok"
		a.Data = "{}"
	}else{
		a.Code= 1001
		a.Msg= "错误的原因"
		a.Data = "{}"
	}
	m,err:=json.Marshal(a)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(m))
}
