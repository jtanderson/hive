package drone

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

func (d *Drone) StartService() bool {
	return true
}

func (d *Drone) StopService() bool {
	return true
}
