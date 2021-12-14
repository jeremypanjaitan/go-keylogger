package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/eiannone/keyboard"
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
			log.Fatal(err)
		}
		fmt.Println(strings.Trim(msg, "\n"))
	}

}

func (service *Service) SendKeyLog() {
	defer service.Wg.Done()
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		service.Conn.Write([]byte(fmt.Sprintf("You pressed: rune %q, key %X\r\n", char, key)))
		if key == keyboard.KeyEsc {
			os.Exit(0)
		}
	}
}
