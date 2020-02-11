package main

import (
	"fmt"
	"time"
)

func main() {
	var keyToCheck uintptr = 0x46  // 0x46 = F key
	fmt.Println(keyIsPressed(keyToCheck))

	x, y := getCursorPos()
	fmt.Println("current cursor pos", x, y)

	for i:=0;i<10;i++ {
		moveMouse(1, 0)
		time.Sleep(1 * time.Millisecond)
	}

	x, y = getCursorPos()
	fmt.Println("current cursor pos", x, y)
}

