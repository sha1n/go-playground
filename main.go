package main

import (
	"./ctrl"
	"./json"
	"./methods"
	"./pkg"
	"./routines"
)

func main() {
	ctrl.Run()
	methods.Run()
	json.Run()
	pkg.Run()
	routines.Run()
}
