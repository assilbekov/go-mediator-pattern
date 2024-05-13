package pkg

import "fmt"

type Passenger struct {
	Dispatcher
}

func (p *Passenger) Arrive() {
	if p.CanArrive(p) {
		fmt.Println("Passenger arrived")
		p.NotifyArrival(p)
		return
	}
	fmt.Println("Passenger denied entry")
}

func (p *Passenger) Depart() {
	fmt.Println("Passenger departed")
	p.Dispatcher.NotifyArrival(p)
}

func (p *Passenger) PermitArrival() {
	fmt.Println("Passenger arrived")
	p.Arrive()
}
