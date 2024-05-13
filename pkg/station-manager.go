package pkg

import "fmt"

type StationManager struct {
	PlatformFree bool
	TrainQueue   []Vehicle
}

func NewStationManager() *StationManager {
	return &StationManager{
		PlatformFree: true,
		TrainQueue:   make([]Vehicle, 0),
	}
}

func (s *StationManager) CanArrive(v Vehicle) bool {
	if s.PlatformFree {
		fmt.Println("Platform is free")
		s.PlatformFree = false
		return true
	}
	fmt.Println("Platform is occupied")
	s.TrainQueue = append(s.TrainQueue, v)
	return false
}

func (s *StationManager) NotifyArrival(v Vehicle) {
	s.PlatformFree = true
	fmt.Println("Platform is free")
	if len(s.TrainQueue) > 0 {
		nextTrain := s.TrainQueue[0]
		s.TrainQueue = s.TrainQueue[1:]
		nextTrain.PermitArrival()
	}
}
