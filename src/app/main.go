package main

import (
	pincontroller "github.com/jfox/piApi/src/controller"
	pinservice "github.com/jfox/piApi/src/service"
)

func main() {

	pinservice.SetUp()
	pincontroller.SetUpAPI()

}
