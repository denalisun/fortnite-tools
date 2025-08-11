package core

import (
	"fmt"

	"github.com/denalisun/fortnite-tools/utilities"
)

type MenuOption struct {
	Name     string
	Id       int
	Callback func(*int) bool
}

type Menu struct {
	Name    string
	Id      int
	Options []*MenuOption
}

var (
	ALL_MENUS        []Menu
	CURRENT_MENU     int = 0
	CURRENT_SELECTED int = 0
)

func RegisterNewMenu(name string, options []*MenuOption) {
	menu := Menu{Name: name, Id: len(ALL_MENUS), Options: options}
	ALL_MENUS = append(ALL_MENUS, menu)
}

func FindMenuByID(id int) *Menu {
	for menuI := range ALL_MENUS {
		Imenu := ALL_MENUS[menuI]
		if Imenu.Id == id {
			return &Imenu
		}
	}
	return nil
}

func ChangeMenu(id int) {
	CURRENT_MENU = id
	CURRENT_SELECTED = 0
}

func PrintCurrentMenu() error {
	for menuI := range ALL_MENUS {
		Imenu := ALL_MENUS[menuI]
		if Imenu.Id == CURRENT_MENU {
			Imenu.PrintMenu()
		}
	}
	return fmt.Errorf("Couldn't find current menu!")
}

func HandleControls() {
	isHoldingDownArrow := false
	isHoldingUpArrow := false
	isHoldingSpaceOrEnter := false

	var m *Menu

	for {
		m = &ALL_MENUS[CURRENT_MENU]
		if utilities.GetKeyDown(0x28) {
			if !isHoldingDownArrow {
				isHoldingDownArrow = true

				CURRENT_SELECTED += 1
				if CURRENT_SELECTED > (len(m.Options) - 1) {
					CURRENT_SELECTED = 0
				}

				m.PrintMenu()
			}
		} else {
			isHoldingDownArrow = false
		}

		if utilities.GetKeyDown(0x26) {
			if !isHoldingUpArrow {
				isHoldingUpArrow = true

				CURRENT_SELECTED -= 1
				if CURRENT_SELECTED < 0 {
					CURRENT_SELECTED = (len(m.Options) - 1)
				}

				m.PrintMenu()
			}
		} else {
			isHoldingUpArrow = false
		}

		// Selecting
		if utilities.GetKeyDown(0x20) {
			if !isHoldingSpaceOrEnter {
				isHoldingSpaceOrEnter = true
				if m.Options[CURRENT_SELECTED].Callback != nil {
					shouldRefresh := m.Options[CURRENT_SELECTED].Callback(&CURRENT_SELECTED)
					if shouldRefresh {
						PrintCurrentMenu()
					}
				}
			}
		} else {
			isHoldingSpaceOrEnter = false
		}
	}
}

func (m Menu) PrintMenu() {
	utilities.ClearScreen()

	width, height := utilities.GetTerminalSize()
	centerWidth := width / 2

	utilities.PrintfToLocation(int(centerWidth)-(len(m.Name)/2), 2, m.Name)

	// Print all options
	for optI := range m.Options {
		option := m.Options[optI]
		if CURRENT_SELECTED == optI {
			nme := fmt.Sprintf("> %s <", option.Name)
			utilities.PrintfToLocation(int(centerWidth)-(len(nme)/2), 4+optI, nme)
		} else {
			utilities.PrintfToLocation(int(centerWidth)-(len(option.Name)/2), 4+optI, option.Name)
		}
	}

	utilities.PrintfToLocation(0, int(height), "Made by denalisun (2025)")
	utilities.PrintfToLocation(int(width)-29, int(height), "Press ENTER or SPACE to select")
}
