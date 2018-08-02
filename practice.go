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
//add do method to get data (HMGET)
func submituser(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w,"1")
	var a response
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
	fmt.Println(err)
	}
	exists ,err := redis.Bool(c.Do("HGET","submituser","1"))
	if err != nil{
		fmt.Println("error", err)
	}
	//c.Send("HMGET", "submituser","")
	//c.Flush()
	//c.Receive() // reply from SET
	//v, err = c.Receive() // reply from GET
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
	fmt.Fprintf(w,string(m))
	}


func queryuser (w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "1")
	var a response
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	exists ,err := redis.Bool(c.Do("HEXISTS","submituser","1"))
	if err != nil{
		fmt.Println(err)
	}

	if exists {
		c.Send("HGET", "submituser","1")
		c.Flush()
		v, err := c.Receive() // reply from GET
		if err != nil{
			fmt.Println(err)
		}
		var x string
		a.Code = 0
		a.Msg = "ok"
		a.Data= fmt.Sprintf(x,v)
	}else{
		a.Code= 1001
		a.Msg= "错误的原因"
		a.Data= "{}"
	}
	m,err:=json.Marshal(a)
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Fprintln(w,string(m))

}
