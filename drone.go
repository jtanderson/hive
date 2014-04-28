package main

import (
	"fmt"
	"github.com/jtanderson/hive/drone"
)

func main() {
	d := drone.Drone{}
	fmt.Println("Starting drone", d.Id)
}
