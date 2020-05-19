package main

import (
	"fmt"
)

type Prototype struct {
	Value int
}

func (p *Prototype) Clone() *Prototype {
	new_p := *p
	return &new_p
}

func main() {
	p := &Prototype{1}

	fmt.Printf("p value: %d\n", p.Value)

	p1 := p.Clone()
	p2 := p
	p.Value = 100

	fmt.Printf("p value: %d\n", p.Value)
	fmt.Printf("p1 value: %d\n", p1.Value)
	fmt.Printf("p2 value: %d\n", p2.Value)
}