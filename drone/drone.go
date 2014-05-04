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

type Drone struct {
	Id int
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

func (d *Drone) MyCall(a int, b *int) error {
	/*	err := errors.New("test")*/
	/*	return err*/
	return nil
}
