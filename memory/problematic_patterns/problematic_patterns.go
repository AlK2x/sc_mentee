package main

import (
	"math/rand"
)

type Foo struct {
	value string
	num   int
	bar   *Bar
}

type Bar struct {
	value string
	num   int
	foo   *Foo
}

type ConstantGrow struct {
	m map[int]*Foo
	s []*Bar
}

func NewConstantGrow() *ConstantGrow {
	return &ConstantGrow{
		m: make(map[int]*Foo),
		s: make([]*Bar, 0),
	}
}

func (cg *ConstantGrow) Grow(n int) {
	startFrom := len(cg.s)
	for i := range n {
		number := startFrom + i
		f := &Foo{
			value: "foo",
			num:   number,
			bar: &Bar{
				value: "bar",
				num:   number,
			},
		}
		f.bar.foo = f
		cg.s = append(cg.s, f.bar)
		cg.m[number] = f
	}
}

func rundomNumberGenerator() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for range 1000 {
			ch <- rand.Intn(100)
		}
	}()
	return ch
}

func main() {
	constantGrow := NewConstantGrow()
	for val := range rundomNumberGenerator() {
		constantGrow.Grow(val)
	}
}
