package goreacto

type Processor[T any, R any] interface {
	execute(T) R
}

type MapProcessor[T any, R any] struct {
	block func(T) R
}

func (p *MapProcessor[T, R]) execute(value T) R {
	return p.block(value)
}

func Map[T, R any](p *Publisher[T], block func(T) R) *Publisher[R] {
	new_pub := Publisher[R]{}
	p.Subscribe(func(t T) {
		new_pub.Publish(block(t))
	})

	return &new_pub
}

type FilterProcessor[T any] struct {
	block func(T) bool
}

func (p *FilterProcessor[T]) execute(value T) *T {
	if p.block(value) {
		return &value
	}
	return nil
}

func (p *Publisher[T]) Filter(filterblock func(T) bool) *Publisher[T] {
	new_pub := Publisher[T]{}
	p.Subscribe(func(t T) {
		if filterblock(t) {
			new_pub.Publish(t)
		}
	})
	return &new_pub
}
