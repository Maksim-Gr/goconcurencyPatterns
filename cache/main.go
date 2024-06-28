package main

import "sync"

type Data struct {
	ID      string
	Payload string
}

type Cache struct {
	mu sync.Mutex
	m  map[string]*Data
}

var numCalls int
var numCallsLock sync.Mutex

func retrieveData(ID string) (*Data, bool) {
	numCallsLock.Lock()
	defer numCallsLock.Unlock()
	numCalls++
	return &Data{
		ID:      ID,
		Payload: "payload",
	}, true
}

func (c *Cache) Get(ID string) (Data, bool) {
	c.mu.Lock()
	data, exists := c.m[ID]
	c.mu.Unlock()
	if exists {
		if data == nil {
			return Data{}, false
		}
		return *data, true
	}

	data, loaded = retrieveData(ID)
}
