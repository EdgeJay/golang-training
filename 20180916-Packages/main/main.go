package main

import (
	"github.com/edgejay/golang-training/20180916-Packages/foo"
)

func main() {
	car := &foo.Car{
		NumOfDoors:  5,
		NumOfWheels: 4,
		Engine: &foo.Engine{
			Type: foo.V8,
		},
	}

	car.Info()
	car.Ignition()
}
