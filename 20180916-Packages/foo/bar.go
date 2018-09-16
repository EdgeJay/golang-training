package foo

import "fmt"

const V8 string = "V8"

type Engine struct {
	Type string
}

type Car struct {
	NumOfDoors  int
	NumOfWheels int
	Engine      *Engine
}

func (car *Car) Info() {
	fmt.Printf("%T I have %d door(s), %d wheel(s)\n", car, car.NumOfDoors, car.NumOfWheels)
	fmt.Printf("%T I have a %s engine\n", car.Engine, car.Engine.Type)
}

func (car *Car) Ignition() {
	fmt.Println("Vrooom!")
}
