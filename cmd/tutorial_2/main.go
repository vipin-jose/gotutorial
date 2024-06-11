package main

import "fmt"

type petrolEngine struct {
	// Fields
	kmpl   uint8
	liters uint8
	//ownerInfo owner
	owner
}

type electricEngine struct {
	// Fields
	kwh   uint8
	kwhpl uint8
	//ownerInfo owner
	owner
}

type engine interface {
	refuel(uint8)
	drive(uint8)
	kmLeft() uint8
}

type owner struct {
	// Fields
	name string
}

func (e petrolEngine) refuel(liters uint8) {
	e.liters += liters
}

func (e electricEngine) refuel(kwh uint8) {
	e.kwh += kwh
}

func (e petrolEngine) drive(km uint8) {
	e.liters -= km / e.kmpl
}

func (e electricEngine) drive(km uint8) {
	e.kwh -= km / e.kwhpl
}

func (e petrolEngine) kmLeft() uint8 {
	return e.kmpl * e.liters
}

func (e electricEngine) kmLeft() uint8 {
	return e.kwhpl * e.kwh
}

func canMakeIt(e engine, km uint8) bool {
	return e.kmLeft() >= km
}

func main() {
	var myEngine petrolEngine = petrolEngine{25, 15, owner{"John Doe"}}
	var electricEngine electricEngine = electricEngine{10, 100, owner{"Jane Doe"}}
	//fmt.Println(myEngine.kmpl, myEngine.liters, myEngine.ownerInfo.name)
	// We can directly access the fields of the owner struct using the petrolEngine struct
	fmt.Println(myEngine.kmpl, myEngine.liters, myEngine.name)

	// Anonymous struct
	var myEngine2 = struct {
		kmpl   uint8
		liters uint8
		owner
	}{25, 15, owner{"John Doe"}}
	fmt.Println(myEngine2.kmpl, myEngine2.liters, myEngine2.name)
	fmt.Println(myEngine.kmLeft())
	fmt.Println(canMakeIt(myEngine, 100))
	fmt.Println(canMakeIt(electricEngine, 150))
}
