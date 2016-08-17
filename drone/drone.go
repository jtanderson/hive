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
	id    string
	queen Queen
}

type Args struct {
	A, B int
}

func (d *Drone) Status(args *Args, reply *string) error {
	return nil
}

func StartService(d *Drone) (net.Listener, error) {
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

func StopService(d *Drone) bool {
	return true
}

func (d *Drone) MyCall(args *Args, reply *string) error {
	log.Println("Calling MyCall with Args:", args)
	*reply = "Success!"
	return nil
}
