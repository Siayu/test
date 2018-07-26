package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"

	"strings"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "{\"name\":\"xiaomi\",\"age\":8}")
}

func main() {
	http.HandleFunc("/test/json", handler)
	http.HandleFunc("/test/getjson", handler2)
	log.Fatal(http.ListenAndServe(":8888", nil))
	}

func handler2(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8888/test/json")
	if err != nil{
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Fprint(w, string(content))
	}
}

