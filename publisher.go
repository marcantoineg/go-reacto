package goreacto

type subfunc[T any] func(T)
type errorfunc func(error)

type Publisher[T any] struct {
	_value      *T
	_err        error
	_subs       []subfunc[T]
	_error_subs []errorfunc
	_done       bool
}

func (p *Publisher[T]) Subscribe(sub_block func(T)) *Publisher[T] {
	p._subs = append(p._subs, sub_block)
	return p
}

func (p *Publisher[T]) OnError(on_error_block func(error)) *Publisher[T] {
	p._error_subs = append(p._error_subs, on_error_block)
	return p
}

func (p *Publisher[T]) Publish(new_value T) {
	if !p._done {
		p._value = &new_value
		for _, sub_func := range p._subs {
			sub_func(*p._value)
		}
	}
}

func (p *Publisher[T]) Error(e error) {
	p._err = e
	p._done = true
	p._subs = nil
	for _, errorBlock := range p._error_subs {
		errorBlock(p._err)
	}
}
