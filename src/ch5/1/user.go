package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func (u user) errChangeEmail(email string) {
	u.email = email
}

type USER user

func (u USER) notify() {
	fmt.Printf("Sending USER Email To %s<%s>\n", u.name, u.email)
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// 接受者若不是指针，那么执行方法的对象是原对象的一个拷贝，所以此处没有更改本身
	lisa.errChangeEmail("bill@newdomain.com")
	lisa.notify()

	// 别名只有数据，方法不会
	xqhuang := USER{"xqhuang", "xqhuang@email.com"}
	xqhuang.notify()
}
