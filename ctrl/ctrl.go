package ctrl

import (
	"fmt"
	"math/rand"
	"time"
)

func Run() {

	fmt.Println("*** Basics / Control")
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)

	i := 1
	for i <= 2 {
		n := rand.Int31n(100)
		if n%2 == 0 {
			fmt.Println("Even!", n)
		} else {
			fmt.Println("Odd!", n)
		}
		i += 1
	}
}
