package main

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L. -lraylib -lGL -lm -lpthread -ldl -lrt -lX11
#include "raylib.h"
*/
import "C"

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	C.InitWindow(screenWidth, screenHeight, C.CString("raylib [textures] example - sprite button"))
	defer C.CloseWindow()

	C.InitAudioDevice()
	defer C.CloseAudioDevice()

	fxButton := C.LoadSound(C.CString("resources/buttonfx.wav"))
	defer C.UnloadSound(fxButton)

	button := C.LoadTexture(C.CString("resources/button.png"))
	defer C.UnloadTexture(button)

	frameHeight := float32(button.height / 3)
	sourceRec := C.struct_Rectangle{0, 0, C.float(button.width), C.float(frameHeight)}

	btnBounds := C.struct_Rectangle{
		C.float(screenWidth)/2 - C.float(button.width)/2,
		C.float(screenHeight)/2 - C.float(button.height)/6,
		C.float(button.width),
		C.float(frameHeight),
	}

	var btnState = 1
	var btnAction = false

	mousePoint := C.struct_Vector2{0, 0}

	C.SetTargetFPS(60)

	for !C.WindowShouldClose() {
		mousePoint = C.GetMousePosition()
		btnAction = false
		if C.CheckCollisionPointRec(mousePoint, btnBounds) {
			if C.IsMouseButtonDown(C.MOUSE_BUTTON_LEFT) {
				btnState = 2
			} else {
				btnState = 1
			}

			if C.IsMouseButtonReleased(C.MOUSE_BUTTON_LEFT) {
				btnAction = true
			}
		} else {
			btnState = 0
		}

		if btnAction {
			C.PlaySound(fxButton)
		}

		sourceRec.y = C.float(btnState) * C.float(frameHeight)

		C.BeginDrawing()
		C.ClearBackground(C.RAYWHITE)
		C.DrawTextureRec(button, sourceRec, C.struct_Vector2{btnBounds.x, btnBounds.y}, C.WHITE)
		C.DrawFPS(10, 10)
		C.EndDrawing()
	}
}
