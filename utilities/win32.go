package utilities

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"unsafe"
)

type Coord struct {
	X int16
	Y int16
}

type SMALL_RECT struct {
	Left   int16
	Top    int16
	Right  int16
	Bottom int16
}

type CONSOLE_SCREEN_BUFFER_INFO struct {
	dwSize              Coord
	dwCursorPosition    Coord
	wAttributes         uint16
	srWindow            SMALL_RECT
	dwMaximumWindowSize Coord
}

var (
	kernel32                       = syscall.NewLazyDLL("kernel32.dll")
	getConsoleScreenBufferInfoProc = kernel32.NewProc("GetConsoleScreenBufferInfo")

	user32               = syscall.NewLazyDLL("user32.dll")
	getAsyncKeyStateProc = user32.NewProc("GetAsyncKeyState")
	enumWindowsProc      = user32.NewProc("EnumWindows")
	getWindowTextProc    = user32.NewProc("GetWindowTextW")
	setWindowPosProc     = user32.NewProc("SetWindowPos")
	moveWindowProc       = user32.NewProc("MoveWindow")
)

const (
	SWP_NOMOVE     int16 = 0x2
	SWP_NOSIZE     int16 = 1
	SWP_NOZORDER   int16 = 0x4
	SWP_SHOWWINDOW int   = 0x0040
)

func MoveWindow(hwnd syscall.Handle, X int, Y int, nWidth int, nHeight int, bRepaint int) (err error) {
	r1, _, e1 := moveWindowProc.Call(uintptr(hwnd), uintptr(X), uintptr(Y), uintptr(nWidth), uintptr(nHeight), uintptr(bRepaint))
	if r1 == 0 {
		if e1 != nil {
			err = e1
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func enumWindows(enumFunc uintptr, lparam uintptr) (err error) {
	r1, _, e1 := enumWindowsProc.Call(enumFunc, lparam)
	if r1 == 0 {
		if e1 != nil {
			err = e1
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func FindWindow(title string) (syscall.Handle, error) {
	var hwnd syscall.Handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		b := make([]uint16, 200)
		_, err := getWindowText(h, &b[0], int32(len(b)))
		if err != nil {
			// ignore the error
			return 1 // continue enumeration
		}
		if syscall.UTF16ToString(b) == title {
			// note the window
			hwnd = h
			return 0 // stop enumeration
		}
		return 1 // continue enumeration
	})
	enumWindows(cb, 0)
	if hwnd == 0 {
		return 0, fmt.Errorf("No window with title '%s' found", title)
	}
	return hwnd, nil
}

func getWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r1, _, e1 := getWindowTextProc.Call(uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	if r1 == 0 {
		if e1 != nil {
			err = e1
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func getAsyncKeyState(vKey int) int16 {
	r1, _, _ := getAsyncKeyStateProc.Call(uintptr(vKey))
	return int16(r1)
}

func GetKeyDown(vKey int) bool {
	return (int(getAsyncKeyState(vKey)) & 0x8000) > 0
}

func SetWindowPos(hWnd syscall.Handle, hWndInsertAfter syscall.Handle, X int, Y int, cX int, cY int, flags uint) (uintptr, error) {
	ret, _, err := setWindowPosProc.Call(uintptr(hWnd), uintptr(hWndInsertAfter), uintptr(X), uintptr(Y), uintptr(cX), uintptr(cY), uintptr(flags))
	return ret, err
}

func GetTerminalSize() (int16, int16) {
	handle, err := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		fmt.Println(err)
		return -1, -1
	}

	var csbi CONSOLE_SCREEN_BUFFER_INFO
	ret, _, _ := getConsoleScreenBufferInfoProc.Call(uintptr(handle), uintptr(unsafe.Pointer(&csbi)))
	if ret == 0 {
		fmt.Println("Failed to get console info: returned 0")
		return -1, -1
	}

	columns := csbi.srWindow.Right - csbi.srWindow.Left + 1
	rows := csbi.srWindow.Bottom - csbi.srWindow.Top + 1

	return columns, rows
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\x1b[H\x1b[2J")
	}
}

func PrintfToLocation(x int, y int, format string, a ...any) {
	fmt.Printf("\x1b[%d;%df", y, x)
	fmt.Printf(format, a...)
}
