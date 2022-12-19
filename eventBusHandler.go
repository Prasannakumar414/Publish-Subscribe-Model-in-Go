package main

import (
	"strings"
)

type EventBus struct {
	bus         string
	subscribers map[string][]*Subscriber
}

func EventHandler(eventBus *EventBus) {
	if eventBus.bus != "" {
		event := eventBus.bus
		eventBus.bus = ""
		tmessage := strings.Split(event, "::")
		topic, message := tmessage[0], tmessage[1]
		n := len(eventBus.subscribers[topic])
		for i := 0; i < n; i++ {
			eventBus.subscribers[topic][i].show(message)
		}
	}
}
