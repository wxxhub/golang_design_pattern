package main

import (
	"./strategy"
)

func main()  {
	travel1 := strategy.Travel{
		&strategy.BikeTravel{},
	}

	travel1.Travel()

	travel2 := strategy.Travel{
		&strategy.CarTravel{},
	}

	travel2.Travel()

	travel3 := strategy.Travel{
		&strategy.TrainTravel{},
	}

	travel3.Travel()
}
