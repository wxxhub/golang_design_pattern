package observer

import "fmt"

type Oberver interface {
	Update(data string)
}

type MyObserver struct {
	Name string
}

func (m *MyObserver) Update (data string) {
	fmt.Printf("I'm %s, I recive: %s.\n", m.Name, data)
}