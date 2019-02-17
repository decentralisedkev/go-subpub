# go-subpub

## Requirements

    - No channels

    - A publisher should be able to send out the same data to N subscribers

    - A publisher can choose to send to a specific subscriber

    - A subscriber should be able to ignore certain data that the publisher sends without type assertion 

    - A subscriber cannot modify the original item that the publisher sent.

## Usage
```
func main() {

	// Define subscribers
	subA := &Worker{id: 1}
	subB := &Worker{id: 2}
	subC := &Worker{id: 3}

	// Define a publisher
	pub := Publisher{
		subs: make(map[int]subpub.Subscriber),
	}

	// Subscribe to publisher
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
```

Note: 
_Each subscriber gets a *copy* of the data from the publisher. If you intend to have large payloads, high throughput and a large number of subscribers, then this package may not be the most efficient for this._

_A workaround to the subscriber not having mutable rights to the data is to have the worker process the payload, then signal back to the Publisher what modifications should be done. This way is also safer, as only one object is allowed to edit the data._
