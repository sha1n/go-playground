package datastruct

import (
	"container/list"
	"fmt"
)

// Demo - demo function for this module
func Demo() {
	arrayDemo()
	sliceDemo()
	mapDemo()
	listDemo()
}

func arrayDemo() {
	fmt.Println("*** arrays ***")

	var a [2]int
	a[0] = 1
	a[1] = 2

	fmt.Printf("Array length is %d\r\n", len(a))
	for i := range a {
		fmt.Printf("a[%d]=%d\r\n", i, a[i])
	}
}

func sliceDemo() {
	fmt.Println("*** alices ***")

	s := make([]int, 2)
	s[0] = 1
	s[1] = 2

	fmt.Printf("Slice length is %d\r\n", len(s))
	for i := range s {
		fmt.Printf("s[%d]=%d\r\n", i, s[i])
	}
}

func mapDemo() {
	fmt.Println("*** maps ***")

	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2

	fmt.Printf("Map size is %d\r\n", len(m))
	for k := range m {
		fmt.Printf("m[%s]=%d\r\n", k, m[k])
	}
}

func listDemo() {
	fmt.Println("*** lists ***")

	l := list.New()
	l.PushFront("1")
	f := l.Front()
	s := l.InsertAfter("2", f)
	l.InsertBefore("1.5", s)
	l.PushBack("3")
	l.PushFront("0")

	fmt.Printf("List size is %d\r\n", l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("e=%s\r\n", e.Value)
	}
}
