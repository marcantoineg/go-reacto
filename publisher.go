package goreacto

import (
	"errors"
)

type subfunc[T any] func(T)
type errorfunc[T any] func(T)

type Publisher[T any] struct {
	_value      *T
	_err        error
	_subs       []subfunc[T]
	_error_subs []errorfunc[error]
}

func (p *Publisher[T]) Subscribe(sub_block func(T)) *Publisher[T] {
	p._subs = append(p._subs, sub_block)
	return p
}

func (p *Publisher[T]) OnError(on_error_block func(error)) *Publisher[T] {
	p._error_subs = append(p._error_subs, on_error_block)
	return p
}

func (p *Publisher[T]) onError() {
	if r := recover(); r != nil {
		var err error
		if str, ok := r.(string); ok {
			err = errors.New(str)
		} else if e, ok := r.(error); ok {
			err = e
		}
		p._err = err
		for _, on_error_block := range p._error_subs {
			on_error_block(p._err)
		}
	}
}

func (p *Publisher[T]) Publish(new_value T) {
	p._value = &new_value
	for _, sub_func := range p._subs {
		defer func() { p.onError() }()
		sub_func(*p._value)
	}
}
