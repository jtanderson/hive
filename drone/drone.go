package drone

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type HiveDrone interface {
	Status() string
	StartService() bool
	StopService() bool
}

type Queen struct {
	address     string
	certificate string
}

type Drone struct {
	Id    int
	queen Queen
}

type Args struct {
	A, B int
}

func (d *Drone) Status() bool {
	return true
}

func (d *Drone) StartService() (net.Listener, error) {
	rpc.Register(d)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Println("Starting drone with PID", os.Getpid())
	go http.Serve(listener, nil)
	return listener, e
}

func (d *Drone) StopService() bool {
	return true
}

func (d *Drone) MyCall(args *Args, reply *int) error {
	log.Println("Calling MyCall with Args:", args)
	return nil
}
