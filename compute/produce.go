package compute

func Produce(payload interface{}, ctx *WorkerContext) (ok bool) {
	return !safeProduce(payload, ctx)
}

func TryProduce(payload interface{}, ctx *WorkerContext) (ok bool, rejected interface{}) {
	if cap(ctx.bufferedChannel) == len(ctx.bufferedChannel) {
		return true, payload
	} else {
		return !safeProduce(payload, ctx), nil
	}
}

func safeProduce(payload interface{}, ctx *WorkerContext) (closed bool) {
	defer func() {
		closed = recover() != nil
	}()

	if !closed {
		ctx.bufferedChannel <- payload
	}

	return false
}
