package json

import "encoding/json"
import "fmt"

type song struct {
	Name   string
	Artist string
}

func Demo() {

	fmt.Println("*** json ***")

	s := song{Name: "Fell On Black Days", Artist: "Soundgarden"}
	j, _ := json.Marshal(s)

	js := string(j)
	fmt.Println("JSON Marshalled: ", js)

	us := song{}
	json.Unmarshal(j, &us)
	fmt.Println("JSON Unmarshalled: ", us)
}
