package main

import (
	/*	"fmt"*/
	"github.com/jtanderson/hive/drone"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	d := new(drone.Drone)
	rpc.Register(d)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Println("Starting drone with PID", os.Getpid())

	go http.Serve(listener, nil)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
	log.Println("Stopping drone...")
}
