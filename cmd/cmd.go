package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

// Demo - demo function for this module
func Demo() {
	fmt.Println("*** cmd/pipes ***")

	ls := exec.Command("ls", "-l", "/")
	wc := exec.Command("wc", "-l")

	r, w := io.Pipe()
	ls.Stdout = w
	wc.Stdin = r

	var buffer bytes.Buffer
	wc.Stdout = &buffer

	ls.Start()
	wc.Start()

	ls.Wait()
	w.Close()
	wc.Wait()

	wcOut := strings.TrimSpace(string(buffer.Bytes()))
	fmt.Printf("There are %s files in your root directory\r\n", wcOut)
}
