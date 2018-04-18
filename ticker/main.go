package main

import (
	"fmt"
	"log"
	"time"
)

type Stoppable interface {
	Stop() error
}

type Ticker struct {
	Ticking bool
}

func (t Ticker) Start() {
	t.Ticking = true
	go func() {
		for t.Ticking {
			fmt.Println("tick")
			time.Sleep(1 * time.Second)
		}
	}()
}

func (t Ticker) Stop() error {
	t.Ticking = false
	return nil
}

func StopArbitraryStoppable(s Stoppable) error {
	return s.Stop()
}

func main() {
	t := Ticker{}
	t.Start()
	time.Sleep(10 * time.Second)
	if err := StopArbitraryStoppable(t); err != nil {
		log.Fatalf("error: %v", err)
	}
}
