package compute

import (
	"sync"
	"time"
)

type Worker struct {
	wg      *sync.WaitGroup
	process func(interface{})
}

func NewWorker(process func(interface{})) *Worker {
	return &Worker{
		wg:      &sync.WaitGroup{},
		process: process,
	}
}

func (worker *Worker) AwaitShutdown() {
	worker.wg.Wait()
}

func (worker *Worker) Start(context *WorkerContext) {
	worker.wg.Add(1)

	go func() {
		consume(worker, context)
		worker.wg.Done()
	}()
}

func consume(worker *Worker, context *WorkerContext) {
	for {
		select {
		case <-context.StopSigChan:
			close(context.bufferedChannel)
			return

		case payload := <-context.bufferedChannel:
			worker.process(payload)

		case <-time.After(time.Second):
		}
	}
}
