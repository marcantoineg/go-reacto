package goreacto

import "errors"

type Publisher[T any] struct {
	_value    *T
	_err      error
	_subs     []Subscriber[T]
	_on_error []func(error)
}

func (p *Publisher[T]) Subscribe(sub_block func(T)) *Publisher[T] {
	p._subs = append(p._subs, Subscriber[T]{sub_block})
	return p
}

func (p *Publisher[T]) OnError(on_error_block func(error)) *Publisher[T] {
	p._on_error = append(p._on_error, on_error_block)
	return p
}

func onError[T any](p Publisher[T]) {
	for _, on_error_block := range p._on_error {
		on_error_block(p._err)
	}
}

func (p *Publisher[T]) Publish(new_value T) {
	p._value = &new_value
	for _, sub := range p._subs {
		defer func() {
			if r := recover(); r != nil {
				var err error
				if str, ok := r.(string); ok {
					err = errors.New(str)
				} else if e, ok := r.(error); ok {
					err = e
				}
				p._err = err
				onError(*p)
			}
		}()
		sub.run(*p._value)
	}
}
