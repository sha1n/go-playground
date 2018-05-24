package interf

import (
	"fmt"

	"github.com/fatih/color"
)

type printable interface {
	print()
}

type coloredPrintable interface {
	printGreen()
	printRed()
}

type console interface {
	printable
	coloredPrintable
}

type text string

func (t text) print() {
	fmt.Printf("text: [%s]\r\n", t)
}
func (t text) printGreen() {
	color.Green("text: [%s]", t)
}
func (t text) printRed() {
	color.Red("text: [%s]", t)
}

// Demo - demo function for this module
func Demo() {

	fmt.Println("*** interfaces ***")

	var p console
	s := text("I'm a printable")
	p = s

	p.print()
	p.printGreen()
	p.printRed()
	fmt.Println()
}
