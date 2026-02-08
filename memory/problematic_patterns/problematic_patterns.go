package main

import (
	"math/rand"

	"net/http"
	_ "net/http/pprof"

	"github.com/grafana/pyroscope-go"
)

// cycle references
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

type Stat struct {
	fooCnt int
	barCnt int
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

// constant growing and allocation of small objects
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

// heap alloation, m.b. better return value, not a pointer
func (cg *ConstantGrow) HeapAllocation() *Stat {
	stat := &Stat{}
	stat.barCnt = len(cg.m)
	stat.fooCnt = len(cg.s)
	return stat
}

func rundomNumberGenerator() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for range 100000 {
			ch <- rand.Intn(100)
		}
	}()
	return ch
}

func main() {
	pyroscope.Start(pyroscope.Config{
		ApplicationName: "simple.golang.app",
		ServerAddress:   "http://localhost:8080",
		Logger:          pyroscope.StandardLogger,
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
		},
	})
	constantGrow := NewConstantGrow()
	for val := range rundomNumberGenerator() {
		constantGrow.Grow(val)
		_ = constantGrow.HeapAllocation()
	}

	http.ListenAndServe("localhost:6060", nil)
}
