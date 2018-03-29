package refs

import (
	"fmt"

	"github.com/fatih/color"
)

type data struct {
	value string
}

// Demo - demo function for this module
func Demo() {
	fmt.Println("*** pointers ***")

	a := "a"
	fmt.Printf("a=%s and it's address is 0x%x\r\n", a, &a)

	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	d := data{value: "v1"}
	fmt.Printf("Struct d has value %s\r\n", green(d.value))
	fmt.Printf("'byval' returned struct with value %s\r\n", green(byval(d).value))
	fmt.Printf("But d still has value %s\r\n", green(d.value))
	fmt.Printf("'byref' returned struct with value %s\r\n", green(byref(&d).value))
	fmt.Printf("And now d has changed its value to %s\r\n", red(d.value))
}

func byval(d data) data {
	d.value = "newvalue"
	return d
}

func byref(d *data) *data {
	d.value = "newvalue"
	return d
}
