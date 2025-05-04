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
	//fmt.Println("DONES")
	//x := float32(0)
	//y := float32(0)
	//for ; x <= constants.ScreenWidth; x += 10 {
	//	lines = append(lines, models.Line{
	//		StartX:      positionX,
	//		StartY:      positionY,
	//		EndX:        x,
	//		EndY:        y,
	//		StrokeWidth: 1.0,
	//		Color:       color.White,
	//		Aa:          true,
	//	})
	//}
	//
	//for ; y <= constants.ScreenHeight; y += 10 {
	//	lines = append(lines, models.Line{
	//		StartX:      positionX,
	//		StartY:      positionY,
	//		EndX:        x,
	//		EndY:        y,
	//		StrokeWidth: 1.0,
	//		Color:       color.White,
	//		Aa:          true,
	//	})
	//}
	//
	//for ; x >= 0; x -= 10 {
	//	lines = append(lines, models.Line{
	//		StartX:      positionX,
	//		StartY:      positionY,
	//		EndX:        x,
	//		EndY:        y,
	//		StrokeWidth: 1.0,
	//		Color:       color.White,
	//		Aa:          true,
	//	})
	//}
	//
	//for ; y >= 0; y -= 10 {
	//	lines = append(lines, models.Line{
	//		StartX:      positionX,
	//		StartY:      positionY,
	//		EndX:        x,
	//		EndY:        y,
	//		StrokeWidth: 1.0,
	//		Color:       color.White,
	//		Aa:          true,
	//	})
	//}
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
}

func (g *Game) Update() error {
	if frameTimer >= 60 {
		secTimer++
		frameTimer = 0
		//circleColor = randomColor()
	} else {
		//Ball moving logic
		if positionX >= constants.ScreenWidth-constants.LightSourceRadius-1 ||
			positionX <= 0+constants.LightSourceRadius {
			speedX = speedX * -1
		}
		if positionY >= constants.ScreenHeight-constants.LightSourceRadius-1 ||
			positionY <= 0+constants.LightSourceRadius {
			speedY = speedY * -1
		}
		positionX += float32(speedX)
		positionY += float32(speedY)
		for i := range lines {
			changePositionOfLine(speedX, speedY, &lines[i])
		}
		frameTimer++
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	vector.DrawFilledCircle(screen, positionX, positionY, radius, circleColor, true)
	for _, v := range lines {
		vector.StrokeLine(screen, positionX, positionY, v.EndX, v.EndY, v.StrokeWidth, v.Color, v.Aa)
	}
	//ebitenutil.DebugPrint(screen, "Seconds passed: "+strconv.Itoa(secTimer))

}

//func (g *Game) Draw(screen *ebiten.Image) {

//vector.DrawFilledCircle(screen, positionX, positionY, radius, circleColor, true)
//for _, v := range lines {
//sx := v.StartX + float32(speedX)
//ex := v.EndX + float32(speedX)
//sy := v.StartY + float32(speedY)
//ey := v.EndY + float32(speedY)
//vector.StrokeLine(screen, sx, sy, ex, ey, v.StrokeWidth, v.Color, v.Aa)
//}
//ebitenutil.DebugPrint(screen, "Seconds passed: "+strconv.Itoa(secTimer))

// }
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
