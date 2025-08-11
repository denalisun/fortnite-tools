package main

import (
	"os"

	"github.com/denalisun/fortnite-tools/core"
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

func makeFOVMenu() {
	fovMenu80FOV := core.MenuOption{Name: "80 FOV", Id: 1}
	fovMenu90FOV := core.MenuOption{Name: "90 FOV", Id: 2}
	fovMenu100FOV := core.MenuOption{Name: "100 FOV", Id: 3}
	fovMenu110FOV := core.MenuOption{Name: "110 FOV", Id: 4}
	fovMenu120FOV := core.MenuOption{Name: "120 FOV", Id: 5}
	fovMenuExit := core.MenuOption{Name: "Exit", Id: 6}
	fovMenuExit.Callback = func(selectedOption *int) bool {
		core.ChangeMenu(0)
		return true
	}

	core.RegisterNewMenu("FOV Changer", []*core.MenuOption{
		&fovMenu80FOV,
		&fovMenu90FOV,
		&fovMenu100FOV,
		&fovMenu110FOV,
		&fovMenu120FOV,
		&fovMenuExit,
	})
}

func main() {
	makeMainMenu()
	makeFOVMenu()

	core.PrintCurrentMenu()
	core.HandleControls()
}
