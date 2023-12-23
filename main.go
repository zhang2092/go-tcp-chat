package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	linstener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	defer linstener.Close()
	log.Println("started server on :8888")

	for {
		conn, err := linstener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %v", err)
			continue
		}

		go s.newClient(conn)
	}
}
