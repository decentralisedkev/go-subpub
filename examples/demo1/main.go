package main

// Shows a publisher sending information to multiple subscribers
// And subscribers printing off type and their id number

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

	// dummy payload
	pl := subpub.Payload{
		ID:   2,
		Item: 5,
	}

	// broadcast to all subscribers
	pub.Publish(pl)

}
