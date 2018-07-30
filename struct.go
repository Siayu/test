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
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func handler1 (w http.ResponseWriter, r *http.Request) {
	var name [2] User
	name[0] = User {"Jane", 20}
	name[1] = User {"Bob", 20}
	//var u Userslice
	//u.Users = append(u.Users, User{Name: "Fred", Age: 25})
	//u.Users = append(u.Users, User{Name: "Ann", Age: 40})
	//u.Users = append(u.Users, User{Name: "Sean", Age: 50})
	k:=0
	 for k<=1{
	 f, err := json.Marshal(name[k])
		if err != nil {
			fmt.Println("json err:", err)
		}
		k++
		fmt.Println(string(f))
		fmt.Fprintln(w, string(f))
	}
	//b, err :=json.Marshal(u)
	//if(err!=nil){
	//	fmt.Println("json err:",err)
	//}
	//fmt.Fprintln(w,string(b))
}


func handler2 (w http.ResponseWriter, r *http.Request) {
	var u interface{}
	resp, err := http.Get("http://localhost:9999/test/jsonencode")
	if err != nil {
		//fmt.Println(err)
	} else {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("%s", err)
		}
		//fmt.Println(content)
		//fmt.Fprintln(w, string(content))
		//fmt.Println(string(content))
		k := 0
		for k <= 1 {
			str := content
			json.Unmarshal([]byte(str), &u)
			//fmt.Println(u)
			//fmt.Fprintln(w, u)
			k++
		}
	}
}