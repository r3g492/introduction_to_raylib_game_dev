package main

import rl "github.com/gen2brain/raylib-go/raylib"

/*
red dot 추가
*/
func main() {
	rl.InitWindow(800, 450, "2번째 예시")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.DrawCircle(300, 300, 30, rl.Red)

		rl.EndDrawing()
	}
}
