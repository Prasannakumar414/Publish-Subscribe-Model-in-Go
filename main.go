package main

import "fmt"

func main() {
	eventBus := EventBus{bus: "", subscribers: make(map[string][]*Subscriber)}
	for {
		fmt.Println(" enter 1 to publish\n enter 2 to create subscriber\n enter 3 to Subscribe\n enter 4 to unSubscribe\n enter 5 to exit")
		n := 0
		fmt.Scanln(&n)
		switch n {
		case 1:
			fmt.Println("enter topic:")
			topic := ""
			fmt.Scanln(&topic)
			fmt.Println("enter message:")
			message := ""
			fmt.Scanln(&message)
			publisher(&eventBus, topic, message)
		case 2:
			fmt.Println("enter name of subscriber:")
			name, topic := "", ""
			fmt.Scanln(&name)
			fmt.Println("enter topic to subscribe:")
			fmt.Scanln(&topic)
			sub := createSubscriber(name, topic)
			eventBus.subscribers[topic] = append(eventBus.subscribers[topic], sub)
		case 5:
			break
		}

	}
}
