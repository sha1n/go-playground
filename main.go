package main

import (
	"./ctrl"
	"./json"
	"./methods"
	"./pkg"
)

func main() {
	ctrl.Run()
	methods.Run()
	json.Run()
	pkg.Run()
}
