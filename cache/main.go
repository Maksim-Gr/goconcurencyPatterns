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

	data, loaded := retrieveData(ID)
	c.mu.Lock()
	defer c.mu.Unlock()
	d, exists := c.m[data.ID]
	if exists {
		return *d, true
	}
	if !loaded {
		c.m[ID] = nil
		return Data{}, false
	}
	c.m[data.ID] = data
	return *data, true
}
