package models

import (
	"math/rand"
	"time"
)

type Vehicle struct {
	ID       int
	waitChan chan struct{} // channel to be notified
	parkTime time.Duration
}

func NewVehicle(id int) *Vehicle {
	randomDuration := MinParkingTime + time.Duration(rand.Intn(int(MaxParkingTime-MinParkingTime)))
	return &Vehicle{
		ID:       id,
		waitChan: make(chan struct{}),
		parkTime: randomDuration,
	}
}
