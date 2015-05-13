package main

import "github.com/go-martini/martini"

func main() {

	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello World"
	})
	m.RunOnAddr("127.0.0.1:3000")
}
