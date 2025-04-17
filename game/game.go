package game

import (
	"github.com/Caaki/RayTracingWithGo/constants"
	"github.com/Caaki/RayTracingWithGo/models"
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

var speedX = 1
var speedY = 1

var lines []models.Line

type Game struct {
}

func init() {
	x := float32(0)
	y := float32(0)
	for ; x <= constants.ScreenWidth; x += 10 {
		lines = append(lines, models.Line{
			StartX:      positionX,
			StartY:      positionY,
			EndX:        x,
			EndY:        y,
			StrokeWidth: 1.0,
			Color:       color.White,
			Aa:          true,
		})
	}

	for ; y <= constants.ScreenHeight; y += 10 {
		lines = append(lines, models.Line{
			StartX:      positionX,
			StartY:      positionY,
			EndX:        x,
			EndY:        y,
			StrokeWidth: 1.0,
			Color:       color.White,
			Aa:          true,
		})
	}

	for ; x >= 0; x -= 10 {
		lines = append(lines, models.Line{
			StartX:      positionX,
			StartY:      positionY,
			EndX:        x,
			EndY:        y,
			StrokeWidth: 1.0,
			Color:       color.White,
			Aa:          true,
		})
	}

	for ; y >= 0; y -= 10 {
		lines = append(lines, models.Line{
			StartX:      positionX,
			StartY:      positionY,
			EndX:        x,
			EndY:        y,
			StrokeWidth: 1.0,
			Color:       color.White,
			Aa:          true,
		})
	}

}

func (g *Game) Update() error {
	if frameTimer >= 60 {
		secTimer++
		frameTimer = 0
		circleColor = randomColor()
	} else {
		//if positionX >= constants.ScreenWidth-constants.LightSourceRadius-1 ||
		//	positionX <= 0+constants.LightSourceRadius {
		//	speedX = speedX * -1
		//	time.Sleep(2)
		//}
		//if positionY >= constants.ScreenHeight-constants.LightSourceRadius-1 ||
		//	positionY <= 0+constants.LightSourceRadius {
		//	speedY = speedY * -1
		//	time.Sleep(2)
		//}
		//positionX += float32(speedX)
		//positionY += float32(speedY)
		frameTimer++
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	vector.DrawFilledCircle(screen, positionX, positionY, radius, circleColor, true)
	for _, v := range lines {
		vector.StrokeLine(screen, positionX, positionY, v.EndX, v.EndY, v.StrokeWidth, v.Color, v.Aa)
	}
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
