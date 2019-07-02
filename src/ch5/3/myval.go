package main

import "fmt"

type MyVal struct {
	val int
}

func main() {
	a := MyVal{1}
	fmt.Println(a)

	b := a.Add(2)
	fmt.Println(a)
	fmt.Println(b)

	b = (&a).Add(2)
	fmt.Println(a)
	fmt.Println(b)
}

func (val MyVal) Add(gap int) MyVal {
	val.val += gap
	return val
}
