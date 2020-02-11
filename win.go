package main

import (
	"syscall"
	"unsafe"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	procGetKeyState = user32.NewProc("GetKeyState")
	procGetCursorPos = user32.NewProc("GetCursorPos")
	procMouse_event = user32.NewProc("mouse_event")

)

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162805.aspx
type POINT struct {
	X, Y int32
}

func keyIsPressed(key uintptr) bool {
	Val, _, _ := procGetKeyState.Call(key)
	return Val & 0xF > 1
}

func getCursorPos() (int, int) {
	pt := POINT{}
	procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	return int(pt.X), int(pt.Y)
}

func moveMouse(x, y int) {
	procMouse_event.Call(0x0001, uintptr(x), uintptr(y), 0x0, 0x0)
}
