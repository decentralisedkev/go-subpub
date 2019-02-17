package main

// Shows a publisher sending information to all subscribers
// And subscribers using the payload ID to decide which payload to process

import subpub "github.com/decentralisedkev/go-subpub"

func main() {

	// Define subscribers
	subA := &WorkerA{
		Worker{id: 1},
	}
	subB := &WorkerB{
		Worker{id: 2},
	}
	subC := &WorkerC{
		Worker{id: 3},
	}

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
