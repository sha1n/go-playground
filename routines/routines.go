package routines

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Work - executable worker function definition
type Work func()

// WorkItem a stupid work item struct
type WorkItem struct {
	Description string
	Executable  Work
}

// ExecuteInWorker - method for executing work items using go routines
func (workItem *WorkItem) ExecuteInWorker(waitGroup *sync.WaitGroup) {
	waitGroup.Add(1)

	go func() {
		fmt.Printf("\t[async] ~> Starting %s...\r\n", workItem.Description)
		workItem.Executable()
		fmt.Printf("\t[async] ~> %s done!\r\n", workItem.Description)

		defer waitGroup.Done()
	}()

}

// Demo - demo function for this module
func Demo() {

	fmt.Println("*** Go routines ***")
	/*
	 * When SIGINT or SIGTERM is caught write to the quitChannel
	 */
	quitChannel := make(chan os.Signal)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	waitGroup := new(sync.WaitGroup)

	workItem1 := sleepyWorkItemFactory("Work Item #1", "1s")
	workItem1.ExecuteInWorker(waitGroup)

	workItem2 := sleepyWorkItemFactory("Work Item #2", "0s")
	workItem2.ExecuteInWorker(waitGroup)

	fmt.Println("Press Ctrl+C to exit..")

	/*
	 * Wait until we get the quit message
	 */
	<-quitChannel

	waitGroup.Wait()

	fmt.Println("\r\nOk, Bye!")
}

func sleepyWorkerFactory(duration string) Work {
	work := Work(func() {
		d, _ := time.ParseDuration(duration)
		time.Sleep(d)
	})

	return work
}

func sleepyWorkItemFactory(desc string, sleepDuration string) WorkItem {
	exec := sleepyWorkerFactory(sleepDuration)
	workItem := WorkItem{Description: desc, Executable: exec}

	return workItem
}
