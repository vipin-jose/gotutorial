package main

import "fmt"

type petrolEngine struct {
	// Fields
	kmpl   uint8
	liters uint8
	//ownerInfo owner
	owner
}

type owner struct {
	// Fields
	name string
}

func main() {
	var myEngine petrolEngine = petrolEngine{25, 15, owner{"John Doe"}}
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
}
