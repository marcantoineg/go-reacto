package goreacto

type Subscriber[T any] struct {
	_block func(T)
}

func (s Subscriber[T]) run(value T) {
	s._block(value)
}
