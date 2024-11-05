package view

import (
	"image/color"
	"simulation/models"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type View struct {
	parking     *models.Parking
	window      fyne.Window
	parkingGrid *fyne.Container
	waitGrid    *fyne.Container
}

func NewView(parking *models.Parking, window fyne.Window) *View {
	v := &View{
		parking: parking,
		window:  window,
	}
	v.buildUI()
	return v
}

func (v *View) buildUI() {
	v.parkingGrid = v.createParkingGrid()
	v.waitGrid = v.createWaitGrid()

	mainContainer := container.NewVBox(
		canvas.NewText("Estacionamiento Simulator", color.White),
		v.parkingGrid,
		canvas.NewText("Vehículos en espera", color.White),
		v.waitGrid,
	)

	v.window.SetContent(mainContainer)
	v.window.Resize(fyne.NewSize(400, 600))
	v.window.CenterOnScreen()
}

func (v *View) createParkingGrid() *fyne.Container {
	grid := container.New(layout.NewGridLayout(5))
	for i := 0; i < models.MaxParking; i++ {
		rect := canvas.NewRectangle(color.Gray{Y: 0x80})
		rect.SetMinSize(fyne.NewSize(50, 50))
		grid.Add(rect)
	}
	return grid
}

func (v *View) createWaitGrid() *fyne.Container {
	grid := container.New(layout.NewGridLayout(5))
	for i := 0; i < models.MaxWait; i++ {
		rect := canvas.NewRectangle(color.Gray{Y: 0x50})
		rect.SetMinSize(fyne.NewSize(50, 50))
		grid.Add(rect)
	}
	return grid
}

func (v *View) StartSimulation() {
	go v.parking.Run()
	go v.updateUI()
}

func (v *View) updateUI() {
	for {
		time.Sleep(1 * time.Second)
		for i, vehicle := range v.parking.GetParkingStatus() {
			rect := v.parkingGrid.Objects[i].(*canvas.Rectangle)
			if vehicle == nil {
				rect.FillColor = color.Gray{Y: 0x80}
			} else {
				rect.FillColor = color.RGBA{R: 70, G: 175, B: 55, A: 255}
			}
			rect.Refresh()
		}

		waitCars := v.parking.GetWaitCars()
		for i := 0; i < models.MaxWait; i++ {
			rect := v.waitGrid.Objects[i].(*canvas.Rectangle)
			if i < len(waitCars) {
				rect.FillColor = color.RGBA{R: 250, G: 205, B: 32, A: 255}
			} else {
				rect.FillColor = color.Gray{Y: 0x50}
			}
			rect.Refresh()
		}
		v.window.Content().Refresh()
	}
}
