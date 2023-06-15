package main

import "sync"

type counter struct {
	mu sync.Mutex
	n  int
}

func newCounter() *counter {
	return &counter{
		mu: sync.Mutex{},
		n:  0,
	}
}

func (c *counter) Add() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

func (c *counter) Get() int {
	c.mu.Lock()
	n := c.n
	c.mu.Unlock()
	return n
}

func (c *counter) Reset() {
	c.mu.Lock()
	c.n = 0
	c.mu.Unlock()
}
