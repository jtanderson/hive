package main

import (
	/*	"fmt"*/
	"github.com/jtanderson/hive/drone"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	d := new(drone.Drone)

	d.StartService()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
	log.Println("Stopping drone...")
}
