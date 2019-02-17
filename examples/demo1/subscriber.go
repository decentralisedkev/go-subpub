package main

import (
	"fmt"

	subpub "github.com/decentralisedkev/go-subpub"
)

type Worker struct {
	id int
}

func (s *Worker) Process(p subpub.Payload) {
	switch v := p.Item.(type) {
	case int:
		fmt.Printf("Found an int of value %d Worker: %d\n", v, s.id)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

}
func (c *Worker) ProcessErr(e error) {
	fmt.Println("Processing error here")
}
func (s *Worker) ID() int {
	return s.id
}
