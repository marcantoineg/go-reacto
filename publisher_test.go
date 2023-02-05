package goreacto

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Publishers(t *testing.T) {
	type testrun[T any] struct {
		testName          string
		test_manipulation func(*Publisher[T]) any
		checks            func(*Publisher[T], any)
	}

	testRuns := []testrun[any]{
		{
			"a publisher initializes correctly",
			nil,
			func(p *Publisher[any], _ any) {
				assert.Empty(t, p._error_subs)
				assert.Empty(t, p._subs)
				assert.Nil(t, p._err)
				assert.Nil(t, p._value)
				assert.False(t, p._done)
			},
		},
		{
			"a publisher updates it's value when published to",
			func(p *Publisher[any]) any {
				p.Publish(int(0))

				return nil
			},
			func(p *Publisher[any], _ any) {
				assert.Empty(t, p._error_subs)
				assert.Empty(t, p._subs)
				assert.Nil(t, p._err)
				assert.Equal(t, "int", fmt.Sprintf("%T", *p._value))
				assert.Equal(t, 0, *p._value)
			},
		},
		{
			"a publisher updates it's subs when subscribed to",
			func(p *Publisher[any]) any {
				p.Subscribe(func(a any) {})
				return nil
			},
			func(p *Publisher[any], _ any) {
				assert.Empty(t, p._error_subs)
				assert.Equal(t, 1, len(p._subs))
				assert.Nil(t, p._err)
				assert.Nil(t, p._value)
			},
		},
		{
			"a publisher updates it's error_subs when subscribed to",
			func(p *Publisher[any]) any {
				p.OnError(func(e error) {})
				return nil
			},
			func(p *Publisher[any], _ any) {
				assert.Empty(t, p._subs)
				assert.Equal(t, 1, len(p._error_subs))
				assert.Nil(t, p._err)
				assert.Nil(t, p._value)
			},
		},
		{
			"a subscriber receives all new values published to a publisher",
			func(p *Publisher[any]) any {
				results := []any{}
				p.Subscribe(func(a any) {
					results = append(results, a)
				})

				for i := 0; i < 1_000; i++ {
					p.Publish(i)
				}

				return results
			},
			func(p *Publisher[any], a any) {
				result := a.([]any)
				for i := 0; i < 1_000; i++ {
					assert.Equal(t, i, result[i])
				}
			},
		},
		{
			"a subscriber does not receives previously published value prior to subscribing",
			func(p *Publisher[any]) any {
				p.Publish(0)

				var result *int = nil
				p.Subscribe(func(a any) {
					if v, ok := a.(int); ok {
						*result = v
					}
				})

				return result
			},
			func(p *Publisher[any], a any) {
				result := a.(*int)
				assert.Nil(t, result)
			},
		},
		{
			"error block is ran when an error is published",
			func(p *Publisher[any]) any {
				var e error
				p.OnError(func(err error) {
					e = err
				})

				p.Error(errors.New("this is an error"))

				return e
			},
			func(p *Publisher[any], a any) {
				err, ok := a.(error)
				assert.True(t, ok)
				assert.Equal(t, "this is an error", err.Error())
			},
		},
		{
			"publisher should not publish new value when it is done",
			func(p *Publisher[any]) any {
				values := []any{}
				p.Subscribe(func(a any) {
					values = append(values, a)
				})

				p.Publish(1)
				p.Error(errors.New("some error"))
				p.Publish(2)
				return values
			},
			func(p *Publisher[any], a any) {
				values, ok := a.([]any)
				assert.True(t, ok)
				assert.Equal(t, []any{1}, values)
			},
		},
	}

	for _, tr := range testRuns {
		t.Run(tr.testName, func(t *testing.T) {
			p := &Publisher[any]{}

			var results any
			if tr.test_manipulation != nil {
				results = tr.test_manipulation(p)
			}

			if tr.checks != nil {
				tr.checks(p, results)
			}
		})
	}
}

func Test_Subscribe(t *testing.T) {

}
