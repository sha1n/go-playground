package compute

type Computer struct {
	Context *WorkerContext
	worker  *Worker
}

func NewComputer(proc1 func(interface{}), queueSize int) *Computer {
	return &Computer{
		NewWorkerContext(queueSize),
		NewWorker(proc1),
	}
}

func (c *Computer) Start() {
	c.worker.Start(c.Context)
}

func (c *Computer) Shutdown() {
	c.Context.SendStopSignal()
	c.worker.AwaitShutdown()
}

func (c *Computer) Enqueue(workItem interface{}) (ok bool) {
	return Produce(workItem, c.Context)
}

func (c *Computer) TryEnqueue(work interface{}) (ok bool, rejected interface{}) {
	return TryProduce(work, c.Context)
}
