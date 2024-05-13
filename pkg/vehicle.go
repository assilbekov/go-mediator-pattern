package pkg

type Vehicle interface {
	Arrive()
	Depart()
	PermitArrival()
}

type Dispatcher interface {
	CanArrive(Vehicle) bool
	NotifyArrival(Vehicle)
}
