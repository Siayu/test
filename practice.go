package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"encoding/json"
)

type response struct{
	Code int
	Msg string
	Data string
}
type users struct{
	Uid string
	Name string
	Age string
}

func main() {
	http.HandleFunc("/submituser", submituser)
	http.HandleFunc("/queryuser", queryuser)

	log.Fatal(http.ListenAndServe(":8404", nil))
}
//add do method to get data (HMGET)
func submituser(w http.ResponseWriter, r *http.Request){
	var a response
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
	exists ,err := redis.Bool(c.Do("HEXISTS","submituser","uid"+x))
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
	fmt.Fprintf(w,string(m))
	}


func queryuser (w http.ResponseWriter, r *http.Request){
	var a response
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
	fmt.Fprintln(w,string(m))

}
