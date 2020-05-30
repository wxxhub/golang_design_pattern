package strategy

import "fmt"

type Strategy interface {
	Travel()
}

type BikeTravel struct {

}

func (b *BikeTravel) Travel () {
	fmt.Println("Travel by bike.")
}

type CarTravel struct {

}

func (c *CarTravel) Travel () {
	fmt.Println("Travel by car.")
}

type TrainTravel struct {

}

func (t *TrainTravel) Travel () {
	fmt.Println("Travel by train.")
}