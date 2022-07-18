package main

import (
	"fmt"
	"strings"
)

func Mul(a interface{}, b int) interface{} {
	switch aCast := a.(type) {
	case fmt.Stringer:
		return strings.Repeat(aCast.String(), b)
	case string:
		return strings.Repeat(aCast, b)
	case int:
		return aCast * b
	default:
		return nil
	}
}

type Foo struct {
	bar string
}

func (f Foo) String() string {
	return f.bar
}

func NewFoo(bar string) *Foo {
	return &Foo{bar: bar}
}

func main() {
	fmt.Println(Mul(NewFoo("fmt.Stringer interface"), 2))
	fmt.Println(Mul("python-like-string-mul", 3))
	fmt.Println(Mul(3, 7))
}
