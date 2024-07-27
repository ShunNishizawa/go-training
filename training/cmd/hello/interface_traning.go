package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// T型のメソッドM
func (t T) M() {
	println(t.S)
}

func emptyInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}