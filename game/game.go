package game

import (
	"github.com/Caaki/RayTracingWithGo/constants"
	"github.com/Caaki/RayTracingWithGo/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
	"math/rand"
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
	cursor CursorPosition
	circle models.Circle
}

func NewGame() *Game {
	circleColor := color.White
	var positionX float32 = constants.ScreenWidth / 2
	var positionY float32 = constants.ScreenHeight / 2

	radius := float32(constants.LightSourceRadius)
	g := &Game{
		circle: models.Circle{PositionX: positionX, PositionY: positionY, CircleColor: circleColor, Radius: radius},
	}

	const rayCount = 180
	for i := 0; i < rayCount; i++ {
		angle := float64(i) * (2 * math.Pi / float64(rayCount))
		endX := positionX + float32(math.Cos(angle))*1000 // Big enough to go off-screen
		endY := positionY + float32(math.Sin(angle))*1000

		lines = append(lines, models.Line{
			StartX:      positionX,
			StartY:      positionY,
			EndX:        endX,
			EndY:        endY,
			StrokeWidth: 1,
			Color:       color.White,
			Aa:          true,
		})
	}

	return g
}

func (g *Game) Update() error {

	mx, my := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.moveBall(mx, my)
	}
	g.cursor = CursorPosition{
		x: mx,
		y: my,
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	vector.DrawFilledCircle(screen, g.circle.PositionX, g.circle.PositionY, g.circle.Radius, g.circle.CircleColor, true)
	for _, v := range lines {
		vector.StrokeLine(screen, g.circle.PositionX, g.circle.PositionY, v.EndX, v.EndY, v.StrokeWidth, v.Color, v.Aa)
	}

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

func changePositionOfLine(x, y int, line *models.Line) {
	line.StartX += float32(x)
	line.EndX += float32(x)

	line.StartY += float32(y)
	line.EndY += float32(y)

}

func (g *Game) moveBall(x, y int) {

	if x >= int(g.circle.PositionX)-constants.LightSourceRadius && x <= int(g.circle.PositionX)+constants.LightSourceRadius && y >= int(g.circle.PositionY)-constants.LightSourceRadius && y <= int(g.circle.PositionY)+constants.LightSourceRadius {
		difX := x - g.cursor.x
		difY := y - g.cursor.y
		g.circle.PositionX += float32(difX)
		g.circle.PositionY += float32(difY)

		for i := range lines {
			changePositionOfLine(difX, difY, &lines[i])
		}
	}
}
