package window

import (
	"log"

	"github.com/dfirebaugh/tortuga/config"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	Texture     uint32
	Framebuffer uint32
)

func InitGLFW() *glfw.Window {
	log.Println("GLFW version:", glfw.GetVersionString())
	log.Println("GLFW window classname:", glfw.X11ClassName)
	// glfw.InitHint(glfw.X11ClassName)
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	c := config.Default()
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(
		c.GetScreenWidth()*c.GetScaleFactor(),
		c.GetScreenHeight()*c.GetScaleFactor(),
		c.GetTitle(), nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func InitGL() {
	err := gl.Init()
	if err != nil {
		panic(err)
	}
}

func InitFrameBuffer() {
	{
		gl.GenTextures(1, &Texture)

		gl.BindTexture(gl.TEXTURE_2D, Texture)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.BindImageTexture(0, Texture, 0, false, 0, gl.WRITE_ONLY, gl.RGBA8)
	}

	{
		gl.GenFramebuffers(1, &Framebuffer)
		gl.BindFramebuffer(gl.FRAMEBUFFER, Framebuffer)
		gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, Texture, 0)

		gl.BindFramebuffer(gl.READ_FRAMEBUFFER, Framebuffer)
		gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, 0)
	}
}

func Terminate() {
	glfw.Terminate()
}
