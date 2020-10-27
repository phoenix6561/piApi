package pinservice

import (
	"fmt"

	"os"

	"github.com/stianeikeland/go-rpio"
)

//SetUp ...
func SetUp() {
	fmt.Println("setting up rpio")
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

//TogglePin ...
func TogglePin(pinNum int) {
	pin := rpio.Pin(pinNum)

	pin.Output()

	pin.Toggle()

	fmt.Println(pin.Read())

}
func PrintAllPinStatus() {
	i := 30
	for i < 0 {

		pin := rpio.Pin(i)

		fmt.Println(pin.Read())

		i--
	}

}
