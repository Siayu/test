package main


import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("%q\n", strings.SplitAfterN("a,b,c", "/n", 2))
	fmt.Println(strings.ToLower("ABCDEFGHIJKLMN"))

}
