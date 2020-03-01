package main

import (
	"bpp/protocol"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		protocol.RouterInit()
		wg.Done()
	}()
	protocol.SendQKey(32)
	wg.Wait()
}
