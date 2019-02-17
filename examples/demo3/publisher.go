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

	for _, sub := range p.subs {
		sub.Process(pl)
	}
	return nil
}

func (p *Publisher) SendErr(err error) {
	for _, sub := range p.subs {
		sub.ProcessErr(err)
	}
}
