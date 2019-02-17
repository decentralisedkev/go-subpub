package subpub

// Publisher interface is satisfied by any struct that wants to send values to
// N subscribers
type Publisher interface {
	// Subscribe is called by the Subscriber and
	//informs the Publisher that we would like to subscribe to it's feed
	Subscribe(Subscriber)
	// Unsubscribe is called by the subscriber and
	// informs the publisher to stop sending data to this subscriber
	// the subscriber passes their id as an argument
	Unsubscribe(int)
	// Publish is the function that iterates all functions and sends the Payload to the Subscribers
	Publish() error
}

// Subscriber interface is satisfied by any struct that wants to receive values sent by a Publisher
type Subscriber interface {
	// ID is the unique id of the subscriber. Ideally this should be global, allowing any Publisher to send a value
	// to a specific subscriber
	ID() int
	// Process is invoked by a Publisher with a given payload
	Process(Payload)
	// ProcessErr is invoked by a Publisher,
	//and is called when the Publisher has encountered an error that it wants the subscribers to be aware of
	ProcessErr(error)
}

// Payload struct wraps the item that the Publisher would like to send with an ID
type Payload struct {
	// ID is used so that the subscriber can ignore the package or be notified of what it is before type asserting
	// Note: this assumes that the Publisher and the subscriber have some shared so that they both can derive the same
	// ID for the same struct
	ID int
	// Item is the data that the Publisher would like to send
	Item interface{}
}
