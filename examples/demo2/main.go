package main

// Shows a publisher sending specific information to certain subscribers
// based on their subscriber ID and the payloadID

import subpub "github.com/decentralisedkev/go-subpub"

func main() {

	// Define subscribers
	subA := &Worker{id: 1}
	subB := &Worker{id: 2}
	subC := &Worker{id: 3}

	// Define a publisher
	pub := Publisher{
		subs: make(map[int]subpub.Subscriber),
	}
	pub.Subscribe(subA)
	pub.Subscribe(subB)
	pub.Subscribe(subC)

	pl := subpub.Payload{
		ID:   1,
		Item: 5,
	}
	// broadcast to all subscribers
	pub.Publish(pl)

	pl = subpub.Payload{
		ID:   2,
		Item: "hello",
	}
	// broadcast to all subscribers
	pub.Publish(pl)

	pl = subpub.Payload{
		ID:   3,
		Item: 3.14,
	}
	// broadcast to all subscribers
	pub.Publish(pl)

}
