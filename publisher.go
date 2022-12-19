package main

import (
	"fmt"
)

func publisher(eventBus *EventBus) {
	fmt.Println("enter topic:")
	topic := ""
	fmt.Scanln(&topic)
	fmt.Println("enter message:")
	message := ""
	fmt.Scanln(&message)
	event := (topic + "::" + message)
	eventBus.bus = event
}
