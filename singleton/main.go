package main

import (
	"fmt"
	"sync"
)

type Signleton struct {
	index int
}

var singleton *Signleton = nil
var single_once sync.Once

func (s *Signleton) PrintIndex ()  {
	fmt.Printf("index: %d\n", s.index)
	s.index++
}

func GetInstance() *Signleton {
	single_once.Do(func() {
			singleton = &Signleton{}
		})

	return singleton
}

func main() {
	s := GetInstance()
	s1 := GetInstance()
	s2 := GetInstance()

	s.PrintIndex()
	s1.PrintIndex()
	s2.PrintIndex()
}