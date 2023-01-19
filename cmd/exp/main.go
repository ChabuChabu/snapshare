package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// user := struct {
	// 	Name string
	// }{
	// 	Name: "Chabu Chabu",
	// }
	user := User{
		Name: "Chabu Chabu",
		Age:  34,
		Meta: UserMeta{
			Visits: 4,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
