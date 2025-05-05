package game

import (
	"image/color"
	"math/rand"

	"github.com/Caaki/RayTracingWithGo/constants"
	"github.com/Caaki/RayTracingWithGo/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.ScreenWidth, constants.ScreenHeight
}

func (g *Game) Update() error {

	mx, my := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.moveBall(mx, my)
	}
	g.cursor.x=mx
	g.cursor.y=my

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	vector.DrawFilledCircle(screen, g.circle.PositionX, g.circle.PositionY, g.circle.Radius, g.circle.CircleColor, true)

  vector.DrawFilledRect(screen, g.rectangle.PositionX, g.rectangle.PositionY, float32(g.rectangle.Width), float32(g.rectangle.Height), g.rectangle.Color, true)

	for _, v := range lines {
		vector.StrokeLine(screen, g.circle.PositionX, g.circle.PositionY, v.EndX, v.EndY, v.StrokeWidth, v.Color, v.Aa)
	}

}

func randomColor() color.RGBA {
	return color.RGBA{
		uint8(rand.Intn(256)),
		uint8(rand.Intn(256)),
		uint8(rand.Intn(256)),
		255,
	}
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



func changePositionOfLine(x, y int, line *models.Line) {
	line.StartX += float32(x)
	line.EndX += float32(x)

	line.StartY += float32(y)
	line.EndY += float32(y)

}
