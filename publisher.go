package goreacto

import (
	"errors"
)

type Publisher[T any] struct {
	_value           *T
	_err             error
	_sub_blocks      []func(T)
	_on_error_blocks []func(error)
}

func (p *Publisher[T]) Subscribe(sub_block func(T)) *Publisher[T] {
	p._sub_blocks = append(p._sub_blocks, sub_block)
	return p
}

func (p *Publisher[T]) OnError(on_error_block func(error)) *Publisher[T] {
	p._on_error_blocks = append(p._on_error_blocks, on_error_block)
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
		for _, on_error_block := range p._on_error_blocks {
			on_error_block(p._err)
		}
	}
}

func (p *Publisher[T]) Publish(new_value T) {
	p._value = &new_value
	for _, block := range p._sub_blocks {
		defer func() {
			p.onError()
		}()
		block(*p._value)
	}
}
