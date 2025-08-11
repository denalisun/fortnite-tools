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
)

func getAsyncKeyState(vKey int) int16 {
	r1, _, _ := getAsyncKeyStateProc.Call(uintptr(vKey))
	return int16(r1)
}

func GetKeyDown(vKey int) bool {
	return (int(getAsyncKeyState(vKey)) & 0x8000) > 0
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
