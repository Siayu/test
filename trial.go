package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
)

type User2 struct{
	Name string
	Age int
}

type Userslices struct {
	Users []User2
}

func main(){
	http.HandleFunc("/test/jsonencode", handler4)
	http.HandleFunc("/test/jsondecode", handler5)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func handler4 (w http.ResponseWriter, r *http.Request) {
	var u Userslices
	u.Users = append(u.Users, User2{Name: "Fred", Age: 25})
	u.Users = append(u.Users, User2{Name: "Ann", Age: 40})
	u.Users = append(u.Users, User2{Name: "Sean", Age: 50})
	b, err :=json.Marshal(u)
	if(err!=nil){
		fmt.Println("json err:",err)
	}
	fmt.Fprintln(w,string(b))
}


func handler5 (w http.ResponseWriter, r *http.Request) {
	var u interface{}
	resp, err := http.Get("http://localhost:8888/test/jsonencode")
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("%s", err)
		}
		fmt.Println(content)
		fmt.Fprintln(w, string(content))
		fmt.Println(string(content))
		str := content
		json.Unmarshal([]byte(str), &u)
		fmt.Println(u)
		fmt.Fprintln(w, u)
	}
}
