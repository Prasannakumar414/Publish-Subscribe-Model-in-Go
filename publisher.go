package main

import (
	"fmt"
)

func publisher(eventBus []string) {
	fmt.Println("enter topic:")
	topic := ""
	fmt.Scanln(&topic)
	fmt.Println("enter message:")
	message := ""
	fmt.Scanln(&message)
	eventBus = append(eventBus, topic+"::"+message)
}
