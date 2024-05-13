package pkg

import "fmt"

type Cargo struct {
	Dispatcher
}

func (c *Cargo) Arrive() {
	if c.CanArrive(c) {
		fmt.Println("Cargo arrived")
		c.NotifyArrival(c)
		return
	}
	fmt.Println("Cargo denied entry")
}

func (c *Cargo) Depart() {
	fmt.Println("Cargo departed")
	c.Dispatcher.NotifyArrival(c)
}

func (c *Cargo) PermitArrival() {
	fmt.Println("Cargo arrived")
	c.Arrive()
}
