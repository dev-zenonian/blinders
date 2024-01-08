package suggestion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	A []int
	B []float32
	C []string
}

func TestOption(t *testing.T) {
	s := &TestStruct{}

	var (
		optA = []int{1, 2}
		optB = []float32{0.0, 0.1}
		optC = []string{"A", "B"}
	)
	opts := []Option{
		WithOptionA(optA),
		WithOptionB(optB),
		WithOptionC(optC),
	}
	for _, opt := range opts {
		opt(s)
	}
	assert.Equal(t, optA, s.A)
	assert.Equal(t, optB, s.B)
	assert.Equal(t, optB, s.B)
}

func WithOptionA(a []int) Option {
	return func(i any) {
		switch s := i.(type) {
		case *TestStruct:
			s.A = a
		default:
			fmt.Println("unknow struct")
		}
	}
}

func WithOptionB(b []float32) Option {
	return func(i any) {
		switch s := i.(type) {
		case *TestStruct:
			s.B = b
		default:
			fmt.Println("unknow struct")
		}
	}
}

func WithOptionC(c []string) Option {
	return func(i any) {
		switch s := i.(type) {
		case *TestStruct:
			s.C = c
		default:
			fmt.Println("unknow struct")
		}
	}
}
