package goreacto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testint int

func Test_Processors(t *testing.T) {
	type testrun[T any] struct {
		testName          string
		test_manipulation func(*Publisher[int]) any
		checks            func(*Publisher[int], any)
	}

	testRuns := []testrun[any]{
		{
			"MapProcessor maps values correctly",
			func(p *Publisher[int]) any {
				mappedValues := []int{}

				Map(
					p,
					func(a int) int {
						return a * 2
					},
				).Subscribe(
					func(i int) {
						mappedValues = append(mappedValues, i)
					},
				)

				p.Publish(1)
				p.Publish(2)
				p.Publish(3)

				return mappedValues
			},
			func(p *Publisher[int], values any) {
				v, ok := values.([]int)
				assert.True(t, ok)
				assert.Equal(t, []int{2, 4, 6}, v)
			},
		},
		{
			"FilterProcessor filters values correctly",
			func(p *Publisher[int]) any {
				filteredValues := []int{}

				Filter(
					p,
					func(i int) bool {
						return i%2 == 0
					},
				).Subscribe(
					func(i int) {
						filteredValues = append(filteredValues, i)
					},
				)

				p.Publish(1)
				p.Publish(2)
				p.Publish(3)
				p.Publish(4)
				p.Publish(5)

				return filteredValues
			},
			func(p *Publisher[int], values any) {
				v, ok := values.([]int)
				assert.True(t, ok)
				assert.Equal(t, []int{2, 4}, v)
			},
		},
		{
			"combined processor process publisher correctly",
			func(p *Publisher[int]) any {
				values := []int{}

				mapPub := Map(
					p,
					func(a int) int {
						return a * 2
					},
				)

				Filter(
					mapPub,
					func(i int) bool {
						return i > 100
					},
				).Subscribe(func(i int) {
					values = append(values, i)
				})

				p.Publish(100)
				p.Publish(10)
				p.Publish(25)
				p.Publish(50)
				p.Publish(51)
				p.Publish(1)

				return values
			},
			func(p *Publisher[int], values any) {
				v, ok := values.([]int)
				assert.True(t, ok)
				assert.Equal(t, []int{200, 102}, v)
			},
		},
	}

	for _, tr := range testRuns {
		t.Run(tr.testName, func(t *testing.T) {
			p := &Publisher[int]{}

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
