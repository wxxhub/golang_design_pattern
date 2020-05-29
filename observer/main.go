package main

import (
	"./observer"
	"fmt"
	"sync"
)

var notify_lock sync.Mutex

type Subject struct {
	observers []observer.Oberver
	name 		string
	context 	string
}

func (s *Subject) Attach(o observer.Oberver) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Notify() {
	notify_lock.Lock()
	for _, o := range s.observers {
		o.Update(s.context)
	}
	notify_lock.Unlock()
}

func (s *Subject) ChangeContext (data string)  {
	s.context = data
	s.Notify()
}


func main()  {

	subject := Subject{
		observers: make([]observer.Oberver, 0),
		name: "subject_test",
		context: "",
	}


	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("observer%d", i)
		my_observer := &observer.MyObserver{
			Name: name,
		}
		subject.Attach(my_observer)
	}

	for i := 0; i < 10; i++ {
		data := fmt.Sprintf("test context %d", i)
		subject.ChangeContext(data)
	}
}
