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

func main() {
	http.HandleFunc("/submituser", submituser)
	http.HandleFunc("/queryuser", queryuser)

	log.Fatal(http.ListenAndServe(":8484", nil))
}

func submituser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"1")
	var a response
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
	fmt.Println(err)
	}
	exists ,err := redis.Bool(c.Do("HEXISTS","submituser","name"))
	if err != nil{
		fmt.Println(err)
	}
	if exists {
		a.Code = 0
		a.Msg = "ok"
	}else{
			a.Code= 1001
			a.Msg= "错误的原因"
			}
	m,err:=json.Marshal(a)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Fprintf(w,string(m))
	}


func queryuser (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "1")
	var a response
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	exists ,err := redis.Bool(c.Do("HEXISTS","submituser","uid"))
	if err != nil{
		fmt.Println(err)
	}

	if exists {
		a.Code = 0
		a.Msg = "ok"
	}else{
		a.Code= 1001
		a.Msg= "错误的原因"
	}
	m,err:=json.Marshal(a)
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Fprintln(w,string(m))

}
