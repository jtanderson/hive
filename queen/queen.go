package queen

import (
	"crypto/rand"
	"fmt"
	/*	"github.com/jtanderson/hive/drone"*/
	"io"
	"log"
	"net/rpc"
)

type HiveQueen interface {
	newUUID() (string, error)
}

type Queen struct {
	drones []Drone
}

type Drone struct {
	id     string
	client *rpc.Client
}

func (d Drone) String() string {
	return d.id
}

func (d *Drone) GetId() string {
	return d.id
}

func (d *Drone) GetClient() *rpc.Client {
	return d.client
}

func (queen *Queen) Run() {
	fmt.Println("Starting Queen...")
}

func (queen *Queen) GetDrone(i int) Drone {
	return queen.drones[i]
}

func (queen *Queen) CountDrones() int {
	return len(queen.drones)
}

func (queen *Queen) newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func (queen *Queen) EnlistDrone(address string) error {
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Fatal("dialing:", err)
		return err
	}
	uuid, err := queen.newUUID()
	d := Drone{id: uuid, client: client}
	queen.drones = append(queen.drones, d)
	return nil
}
