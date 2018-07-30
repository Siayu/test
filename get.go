package main

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
)

func main() {
	http.HandleFunc("/test/get",handler6)
	log.Fatal(http.ListenAndServe(":8787", nil))
}

func handler6(w http.ResponseWriter, r *http.Request){
	keys, ok := r.URL.Query()["name"]
	fmt.Println(ok)
	fmt.Println(r.URL.Query()["name"])
	fmt.Println(keys[0])
	if !ok || len(keys[0])<1{
		return
	}
	key := keys[0]
	fmt.Fprintln(w,key)
	u , err := json.Marshal(key)
	if err!=nil{
		fmt.Println(err)
	}
	b := string(u)
	fmt.Fprintln(w,b)

}
