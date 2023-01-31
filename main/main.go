package main

import (
	"fmt"
	goreacto "go-reacto"
)

func main() {
	p := goreacto.Publisher[int]{}

	mapped_p := goreacto.Map(&p, func(i int) string {
		return fmt.Sprintf("block got value: %d", i)
	})

	mapped_p.Subscribe(func(s string) {
		println(s)
	})

	for i := 0; i < 10; i++ {
		p.Publish(i)
	}
}
