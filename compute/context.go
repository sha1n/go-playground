package compute

type StopSignal = struct{}

type WorkerContext struct {
	bufferedChannel chan interface{}
	StopSigChan     chan StopSignal
}

func NewWorkerContext(buffer int) *WorkerContext {
	return &WorkerContext{
		bufferedChannel: make(chan interface{}, buffer),
		StopSigChan:     make(chan StopSignal, 1),
	}
}

func (ctx *WorkerContext) SendStopSignal() {
	go func() {
		ctx.StopSigChan <- struct{}{}
	}()
}
