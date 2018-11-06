package main

import (
	"github.com/sha1n/go-playground/chan"
	"github.com/sha1n/go-playground/cmd"
	"github.com/sha1n/go-playground/compute"
	"github.com/sha1n/go-playground/ctrl"
	"github.com/sha1n/go-playground/datastruct"
	"github.com/sha1n/go-playground/interf"
	"github.com/sha1n/go-playground/json"
	"github.com/sha1n/go-playground/methods"
	"github.com/sha1n/go-playground/pkg"
	"github.com/sha1n/go-playground/refs"
	"github.com/sha1n/go-playground/templates"
)

func main() {
	refs.Demo()
	ctrl.Demo()
	datastruct.Demo()
	methods.Demo()
	json.Demo()
	pkg.Demo()
	cmd.Demo()
	interf.Demo()
	channels.Demo()
	templates.Demo()
	compute.Demo()

	//routines.Demo()
}
