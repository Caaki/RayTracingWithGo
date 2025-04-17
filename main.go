package main

import (
	"fmt"
	game2 "github.com/Caaki/RayTracingWithGo/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &game2.Game{}
	ebiten.SetWindowSize(980, 800)
	ebiten.SetWindowTitle("Ray tracing in go")

	if err := ebiten.RunGame(game); err != nil {
		fmt.Print("Error running app")
		return
	}

}
