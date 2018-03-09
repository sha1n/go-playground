package main

import (
	"./ctrl"
	"./json"
	"./methods"
)

func main() {
	ctrl.Run()
	methods.Run()
	json.Run()
}
