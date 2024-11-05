package models

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	lambda         = 2.0
	MaxWait        = 10
	MaxParking     = 20
	MinParkingTime = 3 * time.Second
	MaxParkingTime = 5 * time.Second
)

type Parking struct {
	waitCars       []*Vehicle
	parking        [MaxParking]*Vehicle
	entranceChan   chan *Vehicle //channel for entrance vehicles
	exitChan       chan *Vehicle //channel for exiting vehicles
	mu             sync.Mutex
	spaceAvailable *sync.Cond
}

func NewParking() *Parking {
	p := &Parking{
		entranceChan: make(chan *Vehicle),
		exitChan:     make(chan *Vehicle),
	}
	p.spaceAvailable = sync.NewCond(&p.mu)
	return p
}

func (p *Parking) Run() {
	go p.handleEntrance()
	go p.handleExit()
	go p.generateVehicles()
}

func (p *Parking) handleEntrance() {
	for car := range p.entranceChan {
		p.mu.Lock()
		if len(p.waitCars) < MaxWait {
			p.waitCars = append(p.waitCars, car)
			fmt.Printf("Vehículo %d añadido a la cola de espera\n", car.ID)
			p.tryToParkVehicle()
		}
		p.mu.Unlock()
	}
}

func (p *Parking) handleExit() {
	for car := range p.exitChan {
		p.mu.Lock()
		p.removeCarFromParking(car)
		fmt.Printf("Vehículo %d salió del estacionamiento\n", car.ID)
		p.spaceAvailable.Broadcast()
		p.mu.Unlock()
	}
}

func (p *Parking) generateVehicles() {
	id := 1
	for {
		interarrivalTime := -math.Log(1-rand.Float64()) / lambda
		time.Sleep(time.Duration(interarrivalTime * float64(time.Second)))
		car := NewVehicle(id)
		id++
		p.entranceChan <- car
	}
}

func (p *Parking) tryToParkVehicle() {
	for len(p.waitCars) > 0 {
		index := p.findEmptyParkingSpace()
		if index != -1 {
			car := p.waitCars[0]
			p.waitCars = p.waitCars[1:]
			p.parking[index] = car
			fmt.Printf("Vehículo %d estacionado en espacio %d\n", car.ID, index)

			go func(c *Vehicle, idx int) {
				time.Sleep(c.parkTime)
				p.exitChan <- c
			}(car, index)
		} else {
			p.spaceAvailable.Wait()
		}
	}
}

func (p *Parking) findEmptyParkingSpace() int {
	for i, spot := range p.parking {
		if spot == nil {
			return i
		}
	}
	return -1
}

func (p *Parking) removeCarFromParking(car *Vehicle) {
	for i, v := range p.parking {
		if v != nil && v.ID == car.ID {
			p.parking[i] = nil
			fmt.Printf("Vehículo %d ha dejado el espacio %d\n", car.ID, i)
			return
		}
	}
}

func (p *Parking) GetParkingStatus() [MaxParking]*Vehicle {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.parking
}

func (p *Parking) GetWaitCars() []*Vehicle {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.waitCars
}
