package main

import (
	"fmt"
	"github.com/jtanderson/hive/queen"
)

func main() {
	q := queen.Queen{}
	fmt.Println("Queen has", q.CountDrones(), "drones.")
	q.Run()
	q.EnlistDrone("127.0.0.1:1234")
	fmt.Println(q.ShowEnlisted())
}
