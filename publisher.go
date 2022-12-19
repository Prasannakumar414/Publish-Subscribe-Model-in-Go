package main

func publisher(eventBus *EventBus, topic string, message string) {
	event := (topic + "::" + message)
	eventBus.bus = event
	EventHandler(eventBus)
}
