package clock

import (
	"math/rand"
	"time"
)

type Clock struct {
	tick *uint
}

// world clock
var tick uint = 1

func New() *Clock {
	c := &Clock{tick: &tick}
	c.init()
	return c
}

func (Clock) init() {
	ticker := time.NewTicker(time.Millisecond)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				rand.Seed(time.Now().UnixNano())
				tick++
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (c Clock) GetTick() uint {
	return *c.tick
}
