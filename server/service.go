package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type Service struct {
	Wg   *sync.WaitGroup
	Conn net.Conn
}

func NewService(Wg *sync.WaitGroup, Conn net.Conn) *Service {
	return &Service{Wg: Wg, Conn: Conn}
}

func (service *Service) ReceiveData() {
	defer service.Wg.Done()

	buffReader := bufio.NewReader(service.Conn)
	for {
		msg, err := buffReader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(strings.Trim(msg, "\n"))
	}

}
