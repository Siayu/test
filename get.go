package main

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
)

type JsonResponse struct{
	Code int64
	Info string
	Data string
}

func main() {
	http.HandleFunc("/test/get",handler6)
	log.Fatal(http.ListenAndServe(":8787", nil))
}

func handler6(w http.ResponseWriter, r *http.Request){
	var a JsonResponse
	keys, ok := r.URL.Query()["name"]
	fmt.Println(ok)
	fmt.Println(r.URL.Query()["name"])
	fmt.Println(keys[0])
	if !ok || len(keys[0])<1{
		a.Code = 1
		a.Info = "error"
	}else {
		a.Code = 0
		a.Info = "ok"
		a.Data = keys[0]
	}
	u , err := json.Marshal(a)
	if err!=nil{
		fmt.Println(err)
	}
	b := string(u)
	fmt.Fprintln(w,b)
}
