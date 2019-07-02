package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) helloworld() {
	fmt.Printf("Hello World, %s\n", u.name)
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user // 基类

	level string
}

func (ad admin) helloworld() {
	fmt.Printf("Hello World, admin[%s]\n", ad.name)
}

func (ad *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", ad.name, ad.email)
}

func main() {
	ad := admin{
		user: user{
			name:  "xqhuang",
			email: "xqhuang@email.com",
		},
		level: "super",
	}

	ad.user.helloworld()
	ad.user.notify()

	// 那怕接受者类型变了，还是调用重写的方法
	ad.helloworld()
	// 调用重写的方法
	ad.notify()
}
