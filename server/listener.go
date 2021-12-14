package main

import (
	"log"
	"net"
	"os"
	"sync"
)

type ServerListener struct {
	Listener net.Listener
	Wg       sync.WaitGroup
}

func NewServerListener() *ServerListener {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return &ServerListener{Listener: listener}
}

func (serverListener *ServerListener) Listen() {
	log.Println("listen for connection....")
	for {
		conn, err := serverListener.Listener.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		serverListener.Wg.Add(1)
		log.Println("incoming connection from ", conn.RemoteAddr().String())
		service := NewService(&serverListener.Wg, conn)
		go service.ReceiveData()
	}

}
