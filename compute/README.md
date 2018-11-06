```go
// Chaining two computers to process messages in a chain (see demo.go)

computer2 := NewComputer(func(message interface{}) {
    fmt.Printf("Computer#2: processing message %v\n", message)
}, 0)

computer1 := NewComputer(func(message interface{}) {
    fmt.Printf("Computer#1: processing message %v\n", message)
    computer2.Enqueue(message)
}, 5)

computer1.Start()
computer2.Start()

produceMessages(10, computer1.Context)

computer1.Shutdown()
computer2.Shutdown()
```