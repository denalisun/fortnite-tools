package main

import (
	"math"
	"os"

	"github.com/denalisun/fortnite-tools/core"
	"github.com/denalisun/fortnite-tools/utilities"
)

func makeMainMenu() {
	titleMenuFOVChangerOption := core.MenuOption{Name: "FOV Changer", Id: 1}
	titleMenuLobbySwapOption := core.MenuOption{Name: "Lobby Swapper", Id: 2}
	titleMenuAutoOptimizer := core.MenuOption{Name: "Auto-Optimizer", Id: 3}
	titleMenuExit := core.MenuOption{Name: "Exit", Id: 4}
	titleMenuExit.Callback = func(selectedOption *int) bool {
		os.Exit(0)
		return false
	}

	titleMenuFOVChangerOption.Callback = func(selectedOption *int) bool {
		core.ChangeMenu(1)
		return true
	}

	core.RegisterNewMenu("Fortnite Tools", []*core.MenuOption{
		&titleMenuFOVChangerOption,
		&titleMenuLobbySwapOption,
		&titleMenuAutoOptimizer,
		&titleMenuExit,
	})
}

func fovResponse(selectedOption *int) bool {
	fov := 80 + (10 * *selectedOption)
	fovRad := utilities.DegreesToRadians(float64(fov))

	win, _ := utilities.GetDesktopWindow()

	// the second number is the vertical FOV of Fortnite
	r := math.Tan(float64(fovRad)/2) / math.Tan(utilities.DegreesToRadians(50.5340158467)/2)
	w, _ := utilities.GetWindowRect(win)
	h := math.Round(float64(w.Right) / r)

	hwnd, _ := utilities.FindWindow("Fortnite  ") // for some reason fortnite has 2 spaces
	if hwnd != 0 {
		center := float64(w.Bottom/2) - (h / 2)
		utilities.MoveWindow(hwnd, 0, int(center), int(w.Right), int(h), 0)
	}

	return false
}

func makeFOVMenu() {
	fovMenu80FOV := core.MenuOption{Name: "80 FOV", Id: 1}
	fovMenu90FOV := core.MenuOption{Name: "90 FOV", Id: 2}
	fovMenu100FOV := core.MenuOption{Name: "100 FOV", Id: 3}
	fovMenu110FOV := core.MenuOption{Name: "110 FOV", Id: 4}
	fovMenu120FOV := core.MenuOption{Name: "120 FOV", Id: 5}
	fovMenu130FOV := core.MenuOption{Name: "130 FOV", Id: 6}
	fovMenuExit := core.MenuOption{Name: "Exit", Id: 7}
	fovMenuExit.Callback = func(selectedOption *int) bool {
		core.ChangeMenu(0)
		return true
	}
	fovMenu80FOV.Callback = fovResponse
	fovMenu90FOV.Callback = fovResponse
	fovMenu100FOV.Callback = fovResponse
	fovMenu110FOV.Callback = fovResponse
	fovMenu120FOV.Callback = fovResponse
	fovMenu130FOV.Callback = fovResponse

	core.RegisterNewMenu("FOV Changer", []*core.MenuOption{
		&fovMenu80FOV,
		&fovMenu90FOV,
		&fovMenu100FOV,
		&fovMenu110FOV,
		&fovMenu120FOV,
		&fovMenu130FOV,
		&fovMenuExit,
	})
}

func main() {
	makeMainMenu()
	makeFOVMenu()

	core.PrintCurrentMenu()
	core.HandleControls()
}
