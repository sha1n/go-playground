package methods

import "fmt"

type item struct {
  id int64
}

func (i *item) print() {
  fmt.Println("Item ID: ", i.id)
}

func Run() {

  fmt.Println("*** methods ***")

  i := item { id: 1234 }
  i.print()
}
