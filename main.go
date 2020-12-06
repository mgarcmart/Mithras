package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	window := initGlfw()
	defer glfw.Terminate()

	window.MakeContextCurrent()
	//program := initGl()
	if err := gl.Init(); err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version ", version)

	window.SetFramebufferSizeCallback(framebufferSizeCallback)
	gl.ClearColor(0.1, 0.1, 0.5, 0.1)

	for !window.ShouldClose() {
		// Here comes the opengl stuff
		processInput(window)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func initGl() uint32 {

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}

func framebufferSizeCallback(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
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

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(640, 480, "Mithra Engine", nil, nil)
	if err != nil {
		panic(err)
	}

	return window
}
