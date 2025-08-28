package main

import (
	"fmt"
	"sync"
)

type ConcurrentPutMap struct {
	mapField   map[int]int
	mutexField sync.Mutex
}

func newCPM() *ConcurrentPutMap {
	return &ConcurrentPutMap{mapField: make(map[int]int)}
}

func (CPMVar *ConcurrentPutMap) put(key, value int) {
	CPMVar.mutexField.Lock()
	CPMVar.mapField[key] = value
	CPMVar.mutexField.Unlock()
}

func main() {

	cpm := newCPM()

	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range 10 {
				cpm.put(j, i)
			}
		}()
	}

	wg.Wait()

	for k, v := range cpm.mapField {
		fmt.Printf("key: %d, value: %d\n", k, v)
	}

}
