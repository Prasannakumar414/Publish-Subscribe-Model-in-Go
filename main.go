package main

func main() {
	eventBus := EventBus{bus: ""}
	for {
		publisher(&eventBus)
		EventHandler(&eventBus)
	}
}
