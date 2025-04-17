package game

import (
	"fmt"
	"github.com/Caaki/RayTracingWithGo/constants"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
	"strconv"
)

var secTimer = 0
var frameTimer = 0
var (
	circleColor color.Color = color.White
	positionX   float32     = constants.ScreenWidth / 2
	positionY   float32     = constants.ScreenHeight / 2
	radius      float32     = constants.LightSourceRadius
)
var speedX = 3
var speedY = 3

type Game struct{}

func (g *Game) Update() error {
	if frameTimer >= 60 {
		secTimer++
		frameTimer = 0
		fmt.Println("One secound passed")
		circleColor = randomColor()
	} else {
		if positionX >= constants.ScreenWidth-constants.LightSourceRadius ||
			positionX <= 0+constants.LightSourceRadius {
			speedX = speedX * -1
		}
		if positionY >= constants.ScreenHeight-constants.LightSourceRadius ||
			positionY <= 0+constants.LightSourceRadius {
			speedY = speedY * -1
		}
		positionX += float32(speedX)
		positionY += float32(speedY)
		frameTimer++
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	vector.DrawFilledCircle(screen, positionX, positionY, radius, circleColor, true)
	ebitenutil.DebugPrint(screen, "Seconds passed: "+strconv.Itoa(secTimer))

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.ScreenWidth, constants.ScreenHeight
}
func randomColor() color.RGBA {
	return color.RGBA{
		uint8(rand.Intn(256)),
		uint8(rand.Intn(256)),
		uint8(rand.Intn(256)),
		255,
	}
}
