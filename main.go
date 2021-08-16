package main

import (
	interfaces "di-ioc/interfaces"
	machine "di-ioc/machine"
	"fmt"
)

func main() {
	typeOfController := "tv"
	myRemoteController := getRemoteController(typeOfController)

	result, err := myRemoteController.TrunOnMachine()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)

}

func getRemoteController(
	typeOfController string) *interfaces.RemoteController {
	if typeOfController == "tv" {
		return interfaces.RemoteController{}.New(injectAirConditional())
	}
	if typeOfController == "airconditional" {
		return interfaces.RemoteController{}.New(injectTv())
	}
	return interfaces.RemoteController{}.New(injectComputer())
}

func injectAirConditional() *machine.AirConditional {
	return machine.AirConditional{}.New()
}

func injectTv() *machine.Tv {
	return machine.Tv{}.New()
}

func injectComputer() *machine.Computer {
	return machine.Computer{}.New()
}
