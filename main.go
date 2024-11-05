package main

import (
	"fmt"
	"simulation/models"
	"simulation/view"

	"fyne.io/fyne/v2/app"
)

func main() {
	fmt.Println("Starting Parking Simulator...")

	myApp := app.New()
	window := myApp.NewWindow("Parkinglot Simulator")
	window.CenterOnScreen()

	parking := models.NewParking()

	v := view.NewView(parking, window)

	v.StartSimulation()

	window.ShowAndRun()
	fmt.Println("Parking Simulator Initialized.")
}
