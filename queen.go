package main

import (
	"fmt"
	"github.com/jtanderson/hive/drone"
	"github.com/jtanderson/hive/queen"
	"log"
)

func main() {
	q := queen.Queen{}
	fmt.Println("Queen has", q.CountDrones(), "drones.")
	q.Run()
	q.EnlistDrone("127.0.0.1:1234")

	d := q.GetDrone(0)
	args := &drone.Args{7, 8}
	var reply int
	err := d.GetClient().Call("Drone.MyCall", args, &reply)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	fmt.Println("Called MyCall to get:", reply)

	fmt.Println(q.ShowEnlisted())
}
