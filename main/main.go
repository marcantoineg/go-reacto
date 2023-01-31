package main

import goreacto "go-reacto"

func main() {
	p := goreacto.Publisher[string]{}

	p.Subscribe(func(s string) {
		println(s)
		panic("panic in the subsribe")
	}).OnError(func(err error) {
		println("an error occured:\n" + err.Error())
	})

	p.Publish("hello world!")
}
