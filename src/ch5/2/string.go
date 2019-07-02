package main

import "fmt"

func main() {
	var s string = "xqhuang"
	var p *string = &s

	fmt.Println(&s)

	*p = "huanglei"

	fmt.Println(&s)
	fmt.Println(p)
}
