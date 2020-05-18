package main

import "fmt"

type Product interface {
	PrintName(name string)
}

type Product1 struct {

}

func (p *Product1) PrintName (name string)  {
	fmt.Printf("Hello, I'm Product1 %s\n", name)
}

type Product2 struct {

}

func (p *Product2) PrintName (name string)  {
	fmt.Printf("Hello, I'm Product2 %s\n", name)
}

type Factory struct {

}

func (f *Factory) CreateProduct(i int) Product {
	if i == 1 {
		return &Product1{}
	} else {
		return &Product2{}
	}

	return nil
}

func main()  {
	var factory Factory
	product1 := factory.CreateProduct(1)
	product2 := factory.CreateProduct(2)

	product1.PrintName("product1")
	product2.PrintName("product2")
}