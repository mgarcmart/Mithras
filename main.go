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
		processInput(window)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func initGl() {

	if err := gl.Init(); err != nil {
		panic(err)
	}
}

func processInput(w *glfw.Window) {
	key := w.GetKey(glfw.KeyEscape)

	switch key {
	case glfw.Press:
		w.SetShouldClose(true)
		break
	}
}
func initGlfw() *glfw.Window {

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err := glfw.CreateWindow(640, 480, "Mithra Engine", nil, nil)
	if err != nil {
		panic(err)
	}

	return window
}
