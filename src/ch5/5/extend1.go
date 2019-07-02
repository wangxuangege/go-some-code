package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user // 基类

	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "xqhuang",
			email: "xqhuang@email.com",
		},
		level: "super",
	}

	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
