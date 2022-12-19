package main

import "fmt"

type Subscriber struct {
	name  string
	topic string
}

func createSubscriber(name string, topic string) *Subscriber {
	return &Subscriber{name: name, topic: topic}
}
func (s *Subscriber) show(message string) {
	fmt.Println("Subscriber " + s.name + " recieved -> " + message)

}
