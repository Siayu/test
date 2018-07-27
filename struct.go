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

type Userslice struct {
	Users []User
}

func main(){
	http.HandleFunc("/test/jsonencode", handler1)
	http.HandleFunc("/test/jsondecode", handler2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler1 (w http.ResponseWriter, r *http.Request){
	var u Userslice
	fmt.Println(User{Name: "Bob", Age:20})
	u.Users = append(u.Users, User{Name: "Fred", Age:25})
	u.Users = append(u.Users, User{Name: "Ann", Age: 40})
	u.Users = append(u.Users, User{Name: "Sean", Age: 50})
	b, err :=json.Marshal(u)
	if err!= nil{
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	fmt.Fprintf(w,string(b))
}


func handler2 (w http.ResponseWriter, r *http.Request) {
	var u Userslice
	resp, err := http.Get("http://localhost:8080/test/jsonencode")
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("%s", err)
		}
		fmt.Fprintln(w, string(content))
		str := string(content)
		json.Unmarshal([]byte(str), &u)
		fmt.Println(u)
		fmt.Fprintln(w, u)
	}
}