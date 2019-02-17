package main

import subpub "github.com/decentralisedkev/go-subpub"

type Publisher struct {
	subs map[int]subpub.Subscriber
}

func (p *Publisher) Subscribe(s subpub.Subscriber) {
	p.subs[s.ID()] = s
}

func (p *Publisher) Unsubscribe(id int) {
	delete(p.subs, id)
}

// Publish is the function that iterates all functions and sends the Payload to the Subscribers
//
func (p *Publisher) Publish(pl subpub.Payload) error {

	for subID, sub := range p.subs {

		// only send payload with ID 2 to first sub
		if pl.ID == 2 && subID == 1 {
			sub.Process(pl)
			continue
		}

		// only send payload with ID 1 to third sub
		if pl.ID == 1 && subID == 3 {
			sub.Process(pl)
			continue
		}

		// only send payload with ID 3 to second sub
		if pl.ID == 3 && subID == 2 {
			sub.Process(pl)
			continue
		}
	}
	return nil
}

func (p *Publisher) SendErr(err error) {
	for _, sub := range p.subs {
		sub.ProcessErr(err)
	}
}
