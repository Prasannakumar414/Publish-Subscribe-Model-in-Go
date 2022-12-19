package main

import (
	"fmt"
	"strings"
)

type EventBus struct {
	bus string
}

func EventHandler(eventBus *EventBus) {
	if eventBus.bus != "" {
		event := eventBus.bus
		eventBus.bus = ""
		tmessage := strings.Split(event, "::")
		topic, message := tmessage[0], tmessage[1]
		fmt.Println(topic)
		if topic == "1" {
			subcriber1(message)
		}
		if topic == "2" {
			subcriber2(message)
		}
		if topic == "3" {
			subcriber3(message)
		}
	}
}
