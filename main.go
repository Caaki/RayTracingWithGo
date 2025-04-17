package main

import (
	"fmt"
	"github.com/Caaki/RayTracingWithGo/constants"
	game2 "github.com/Caaki/RayTracingWithGo/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &game2.Game{}
	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight)
	ebiten.SetWindowTitle("Ray tracing in go")

	if err := ebiten.RunGame(game); err != nil {
		fmt.Print("Error running app")
		return
	}

}
