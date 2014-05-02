package main

import (
	"fmt"
	"github.com/jtanderson/hive/queen"
)

func main() {
	q := queen.Queen{}
	fmt.Println("Queen has", q.CountDrones(), "drones.")
	q.Run()
}
