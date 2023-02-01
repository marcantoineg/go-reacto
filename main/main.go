package main

import (
	"fmt"
	goreacto "go-reacto"
)

func main() {
	p := goreacto.Publisher[int]{}

	p.Subscribe(func(i int) { println(i) })

	mapped_p := goreacto.Map(&p, func(i int) string {
		return fmt.Sprintf("map block got value: %d", i)
	})

	mapped_p.Subscribe(func(s string) { println(s) })

	p.Publish(1)
	p.Publish(2)
	p.Publish(3)
}
