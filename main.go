package main

import (
	"mediator/pkg"
	"time"
)

func main() {
	stationManager := pkg.NewStationManager()
	passengerTrain := &pkg.Passenger{stationManager}
	cargoTrain := &pkg.Cargo{stationManager}

	passengerTrain.Arrive()
	time.Sleep(time.Second * 1)
	cargoTrain.Arrive()
	passengerTrain.Depart()
	cargoTrain.Depart()
	time.Sleep(time.Second * 1)

	passengerTrain.Arrive()
	time.Sleep(time.Second * 1)
	cargoTrain.Arrive()
}
