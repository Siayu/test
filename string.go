package main


import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", "/n", 2))
	fmt.Printf(strings.ToLower("ABCDEFGHIJKLMN"))

}
