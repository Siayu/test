package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"encoding/json"
)

type response1 struct{
	Code int
	Msg string
	Data string
}
type users1 struct {
	response1
	Uid  string
	Name string
	Age  string
}
func main() {
	var a response1
	var b users
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("error is" ,err)
	}
	var x string
	fmt.Println("Input the uid")
	fmt.Scanf("%s", &x)
	exists ,err := redis.Bool(c.Do("HEXISTS","submituser","uid"+x))
	if err != nil{
		fmt.Println("error is ",err)
	}
	if exists {
		c.Do("HGET", "submituser", "uid"+x)
		c.Do("HGET","submituser", "name"+x)
		c.Do("HGET", "submituser", "age"+x)
		v1, err :=redis.String(c.Do("HGET", "submituser", "uid"+x))
		if err != nil {
			fmt.Println(err)
		}
		v2, err :=redis.String(c.Do("HGET","submituser", "name"+x))
		if err != nil {
			fmt.Println(err)
		}
		v3, err :=redis.String(c.Do("HGET", "submituser", "age"+x))
		fmt.Println(v1)
		if err != nil {
			fmt.Println(err)
		}
		b.Uid = v1
		b.Name = v2
		b.Age = v3
		q,err:=json.Marshal(b)
		a.Code = 0
		a.Msg = "ok"
		a.Data = string(q)
	}else{
		a.Code= 1001
		a.Msg= "错误的原因"
		a.Data= "{}"
	}
	m,err:=json.Marshal(a)
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println(string(m))
}