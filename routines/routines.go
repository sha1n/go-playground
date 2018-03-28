package routines

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func executeInWorker(description string, waitGroup *sync.WaitGroup) {
	waitGroup.Add(1)
	defer waitGroup.Done()

	go func() {
		fmt.Println("Executing", description)
	}()

}

// Run executing code in go routine..
func Run() {

	fmt.Println("*** Go routines ***")
	/*
	 * When SIGINT or SIGTERM is caught write to the quitChannel
	 */
	quitChannel := make(chan os.Signal)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	waitGroup := &sync.WaitGroup{}

	executeInWorker("Work Item", waitGroup)

	fmt.Println("Press Ctrl+C to exit..")

	/*
	 * Wait until we get the quit message
	 */
	<-quitChannel

	waitGroup.Wait()

	fmt.Println("\r\nOk, Bye!")
}
