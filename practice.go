package main

import (
	"net/http"
	"github.com/gomodule/redigo/redis"
	"log"
	"fmt"
)

type response struct{
	Code int
	Msg string
	Data string
}

func main() {
	http.HandleFunc("/submituser", submituser)
	http.HandleFunc("/queryuser", queryuser)

	log.Fatal(http.ListenAndServe(":8404", nil))
}

	func err(){
		var a response
		a.Code = 1001
		a.Msg = "err is not nil"
		a.Data = ""
		var b response
		b.Code = 1001
		b.Msg = "参数为空"
		b.Data = ""
		var d response
	}
func submituser(w http.ResponseWriter, r *http.Request) {
	var a response
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("Dial error", err)
		return
	}
	key1, err := r.URL.Query()["uid"]
	key2, err := r.URL.Query()["username"]
	key3, err := r.URL.Query()["age"]
	if err != nil {

		fmt.Fprint(w , )
		return
	}
	if len(key1)==0 || len(key2)==0 || len(key3) ==0 {
		a.Code = 1001
		a.Msg = "参数为空"
		a.Data = ""
		return
	}
	a.Code = 1
	a.Msg = "success"
	a.Data = ""
}

func queryuser (w http.ResponseWriter, r *http.Request){
	c, err := redis.Dial("tcp", ":6379")
	var a response
	if err != nil {
		fmt.Println("Dial error", err)
		return
	}
	key1, err := r.URL.Query()["uid"]
	if err != nil || len(key1[0])==0 {
		a.Code = 1001
		a.Msg = "参数为空"
		a.Data = ""
		return
	}
	
	a.Data = key1[0]
}
