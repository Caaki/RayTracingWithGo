package game

import (
	"image/color"
	"math"

	"github.com/Caaki/RayTracingWithGo/constants"
	"github.com/Caaki/RayTracingWithGo/models"
	"github.com/hajimehoshi/ebiten/v2"
)

var secTimer = 0
var frameTimer = 0
var ()

var speedX = 1
var speedY = 1

var lines []models.Line

type CursorPosition struct {
	x int
	y int
}

type Game struct {
	cursor    CursorPosition
	circle    models.Circle
	rectangle models.Rectangle
}

func NewGame() *Game {
	circleColor := color.White
	var positionX float32 = constants.ScreenWidth / 2
	var positionY float32 = constants.ScreenHeight / 2
	radius := float32(constants.LightSourceRadius)

	mx, my := ebiten.CursorPosition()

	g := &Game{
		circle: models.Circle{PositionX: positionX, PositionY: positionY, CircleColor: circleColor, Radius: radius, Aa: true},
		cursor: CursorPosition{
			x: mx,
			y: my,
		},
		rectangle: models.Rectangle{
			PositionX: 450,
			PositionY: 400,
			Height:    50,
			Width:     100,
			Aa:        true,
			Color:     circleColor,
		},
	}

	const rayCount = 180
	for i := 0; i < rayCount; i++ {
		angle := float64(i) * (2 * math.Pi / float64(rayCount))
		endX := g.circle.PositionX + float32(math.Cos(angle))*1000
		endY := g.circle.PositionY + float32(math.Sin(angle))*1000

		lines = append(lines, models.Line{
			StartX:      g.circle.PositionX,
			StartY:      g.circle.PositionY,
			EndX:        endX,
			EndY:        endY,
			MaxX:        endX,
			MaxY:        endY,
			StrokeWidth: 1,
			Color:       color.White,
			Aa:          true,
		})
	}

	return g
}
