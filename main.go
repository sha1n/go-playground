package main

import (
	"github.com/sha1n/go-playground/cmd"
	"github.com/sha1n/go-playground/ctrl"
	"github.com/sha1n/go-playground/datastruct"
	"github.com/sha1n/go-playground/json"
	"github.com/sha1n/go-playground/methods"
	"github.com/sha1n/go-playground/pkg"
	"github.com/sha1n/go-playground/routines"
)

func main() {
	ctrl.Demo()
	datastruct.Demo()
	methods.Demo()
	json.Demo()
	pkg.Demo()
	cmd.Demo()
	routines.Demo()
}
