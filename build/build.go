package main

import (
	"fmt"
)

type Builder interface {
	selectCPU()
	selectDisk()
	selectMem()
}

type DesktopBuilder struct {

}

func (d *DesktopBuilder)selectCPU()  {
	fmt.Println("select desktop cpu.")
}

func (d *DesktopBuilder)selectDisk()  {
	fmt.Println("select desktop disk.")
}

func (d *DesktopBuilder)selectMem()  {
	fmt.Println("select desktop memory.")
}

type LaptopBuilder struct {

}
func (l *LaptopBuilder)selectCPU()  {
	fmt.Println("select laptop cpu.")
}

func (l *LaptopBuilder)selectDisk()  {
	fmt.Println("select laptop disk.")
}

func (l *LaptopBuilder)selectMem()  {
	fmt.Println("select laptop memory.")
}

type Director struct {
	builder Builder
}


func (d *Director)create()  {
	d.builder.selectCPU()
	d.builder.selectDisk()
	d.builder.selectMem()
}

func main()  {
	//fmt.Println("Hello World!")
	var laptop LaptopBuilder
	directot := Director{&laptop}

	directot.create()
}

