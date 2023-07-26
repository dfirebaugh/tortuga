package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 450
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "block renderer")
	defer rl.CloseWindow()

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(3.0, 3.0, 3.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 60

	rl.SetTargetFPS(120)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraOrbital)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawGrid(10, 1.0)

		rl.DrawCube(rl.NewVector3(0, 0, 0), 1, 1, 1, rl.Red)
		rl.DrawCube(rl.NewVector3(1, 0, 0), .1, .1, .1, rl.Green)

		rl.EndMode3D()

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

}
