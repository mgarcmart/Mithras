package main

import (
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	window := createDisplay(800, 600, "Mithra Engine")
	defer closeDisplay()

	for !window.ShouldClose() {
		clearDisplay(0, 0.15, 0.3, 1.0)
		updateDisplay(window)
	}
}
