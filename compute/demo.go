package compute

import (
	"fmt"
)

func Demo() {
	fmt.Println("*** Compute ***")

	/*
		A computation unit composed of two workers and a producer.
		The producer (the "main" block) sends a message to worker-1, which in turn passes the message to worker-2

		Each worker context can be configured with a different buffer size, to match message production rate and the time it
		takes the corresponding worker to process messages.
	*/

	worker1Context := NewWorkerContext(5)
	worker2Context := NewWorkerContext(0)

	worker1 := NewWorker(func(message interface{}) {
		fmt.Printf("Worker#1: processing message %v\n", message)
		Produce(message, worker2Context)
	})

	worker2 := NewWorker(func(message interface{}) {
		fmt.Printf("Worker#2: processing message %v\n", message)
	})

	worker1.Start(worker1Context)
	worker2.Start(worker2Context)

	produceMessages(10, worker1Context)

	worker1Context.SendStopSignal()
	worker1.AwaitShutdown()

	worker2Context.SendStopSignal()
	worker2.AwaitShutdown()

}

func produceMessages(count int, context *WorkerContext) {
	var i = 1
	for {
		if i == count {
			break
		}

		message := fmt.Sprintf("message-%d", i)

		//Produce(message, context)
		ok, rejected := TryProduce(message, context)

		if !ok {
			fmt.Printf("Channel is closed")
			break
		}

		if rejected != nil {
			fmt.Println(fmt.Errorf("message [%s] has been rejected - worker queue is full", rejected))
		} else {
			i += 1
		}

	}

}
