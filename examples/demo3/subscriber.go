package main

import (
	"fmt"

	subpub "github.com/decentralisedkev/go-subpub"
)

type Worker struct {
	id int
}

func (c *Worker) ProcessErr(e error) {
	fmt.Println("Processing error here")
}
func (s *Worker) ID() int {
	return s.id
}

type WorkerA struct {
	Worker
}

func (s *WorkerA) Process(p subpub.Payload) {

	if p.ID != 2 {
		return
	}

	switch v := p.Item.(type) {
	case int:
		fmt.Printf("In Payload with ID %d , found an int of value %d Worker: %d\n", p.ID, v, s.id)
	case string:
		fmt.Printf("In Payload with ID %d , found a string with content %s Worker: %d\n", p.ID, v, s.id)
	case float32, float64:
		fmt.Printf("In Payload with ID %d , found a float with value %f Worker: %d\n", p.ID, v, s.id)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

}

type WorkerB struct {
	Worker
}

func (s *WorkerB) Process(p subpub.Payload) {

	if p.ID != 1 {
		return
	}

	switch v := p.Item.(type) {
	case int:
		fmt.Printf("In Payload with ID %d , found an int of value %d Worker: %d\n", p.ID, v, s.id)
	case string:
		fmt.Printf("In Payload with ID %d , found a string with content %s Worker: %d\n", p.ID, v, s.id)
	case float32, float64:
		fmt.Printf("In Payload with ID %d , found a float with value %f Worker: %d\n", p.ID, v, s.id)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

}

type WorkerC struct {
	Worker
}

func (s *WorkerC) Process(p subpub.Payload) {

	if p.ID != 3 {
		return
	}

	switch v := p.Item.(type) {
	case int:
		fmt.Printf("In Payload with ID %d , found an int of value %d Worker: %d\n", p.ID, v, s.id)
	case string:
		fmt.Printf("In Payload with ID %d , found a string with content %s Worker: %d\n", p.ID, v, s.id)
	case float32, float64:
		fmt.Printf("In Payload with ID %d , found a float with value %f Worker: %d\n", p.ID, v, s.id)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

}
