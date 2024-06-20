
package main

import (
	"log"
	"movies/api"
	"net"
)

func main(){
	log.Println("Starting listening on port 9090")

	listen, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := api.NewRPCServer()

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}	
}

