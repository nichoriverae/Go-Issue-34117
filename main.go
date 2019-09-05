package main

import (
	"fmt"
	"main/salut"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	s := salut.Salutation{
		Name:  "World",
		Hello: "Hey",
	}

	fmt.Println(s.Hello, s.Name)
}
