package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	window := initGlfw()
	defer glfw.Terminate()

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		// Here comes the opengl stuff
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func initGl() {

	if err := gl.Init(); err != nil {
		panic(err)
	}
}

func initGlfw() *glfw.Window {

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	window, err := glfw.CreateWindow(640, 480, "Mithra Engine", nil, nil)
	if err != nil {
		panic(err)
	}

	return window
}
