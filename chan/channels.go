package channels

import (
	"fmt"
)

type message struct {
	index int
	text  string
}

func interceptor(in chan message) chan message {
	out := make(chan message)

	go func() { // interceptor loop
		for msg := range in {
			msg.text = fmt.Sprintf("InterceptorWrap { %s }", msg.text)
			out <- msg
		}
		close(out)
	}()

	return out
}

func subscribe(in chan message, handler func(msg message)) {
	go func() {
		for msg := range in {
			handler(msg)
		}
	}()
}

func messageHandler(msg message) {
	fmt.Println("Got message:", msg)
}

func generateMessages(out chan message) {
	i := 1
	for i <= 1000 {
		out <- message{
			index: i,
			text:  fmt.Sprintf("Message number %d", i),
		}

		i++
	}
}

func Demo() {
	out := make(chan message)

	interceptedChannel := interceptor(out)
	subscribe(interceptedChannel, messageHandler)

	generateMessages(out)
}
