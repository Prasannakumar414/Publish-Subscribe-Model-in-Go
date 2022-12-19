# Publish-Subscribe-Model-in-Go
 Implemented publish subscribe model in go using web sockets.Pub-Sub model is a communication model where the publisher is not aware of the subscribers and subscribers are not aware of the publishers. The publisher just sends the message and topic of the message. The event bus manages the communication by sending the messages to the subcribers based on the topic they subscribed.

## Publisher:
- Takes topic and message from user and sends it to the event bus.

## Event Bus:
- Maintains a list of sockets for each topic.
- Sends Messages based on the topic to specified subscriber.

## Subscriber:
- Recieves Messages.
- Can subscribe or unsubscribe.

All Communications are done in JSON format.
