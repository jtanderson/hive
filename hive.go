package main

import (
	"fmt"
	"github.com/jtanderson/hive/drone"
	"github.com/jtanderson/hive/queen"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if len(os.Args) == 1 {
		showHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "queen":
		runQueen()
		return
	case "drone":
		runDrone()
		return
	}

	fmt.Println("That command was not recognized.")
	showHelp()
}

func runDrone() {
	d := new(drone.Drone)

	drone.StartService(d)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
	log.Println("Stopping drone...")

	drone.StopService(d)
}

func runQueen() {
	q := new(queen.Queen)

	/* Delete this madness eventually */
	q.EnlistDrone("127.0.0.1:1234")
	d := q.GetDrone(0)
	args := &drone.Args{7, 8}
	var reply string
	err := d.GetClient().Call("Drone.MyCall", args, &reply)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	fmt.Println("Called MyCall to get:", reply)
	/* ------------------- */

	showEnlisted(q)
}

func showEnlisted(q *queen.Queen) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Address", "Status"})

	for _, v := range q.GetDrones() {
		table.Append([]string{v.Id, v.Address, ""})
	}

	table.Render()
}

func showHelp() {
	fmt.Println("Usage:")
	fmt.Println("\thive command [options]")
}
