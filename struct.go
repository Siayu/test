package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
)

type User struct{
	 Name string
	 Age int
}

func main(){
	http.HandleFunc("/test/jsonencode", handler1)
	http.HandleFunc("/test/jsondecode", handler2)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func handler1 (w http.ResponseWriter, r *http.Request) {
	var u [3] User
	u[0] = User{Name: "Fred", Age: 25}
	u[1] = User{Name: "Ann", Age: 40}
	u[2] = User{Name: "Sean", Age: 50}
	 f, err := json.Marshal(u)
		if err != nil {
			fmt.Println("json err:", err)
		}
		fmt.Println(string(f))
		fmt.Fprintln(w, string(f))
}



func handler2 (w http.ResponseWriter, r *http.Request) {
	var u [] User
	resp, err := http.Get("http://localhost:9999/test/jsonencode")
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("%s", err)
		}
		str:= string(content)
		fmt.Fprintln(w,str)
		//fmt.Fprintln(w, content)
		fmt.Fprintln(w, content)
		json.Unmarshal([]byte(str), &u)
		fmt.Println(u[1])
		k :=0
		for k<=2 {
			fmt.Fprintln(w, u[k])
			k++
		}
	}
}
