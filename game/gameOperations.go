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
    g.moveRectangle(mx,my)
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
      changePositionOfLine(difX, difY, &lines[i],g)
    }
  }
}

func (g *Game) moveRectangle (x,y int){
  if x >= int(g.rectangle.PositionX) && x<= int(g.rectangle.PositionX + float32(g.rectangle.Width)) && y>=int(g.rectangle.PositionY) && y <= int(g.rectangle.PositionY) + g.rectangle.Height{
    difX:= x-g.cursor.x
    difY:= y - g.cursor.y
    g.rectangle.PositionX+=float32(difX)
    g.rectangle.PositionY+=float32(difY)
    for i := range lines {
      changePositionOfLineRec(difX, difY, &lines[i],g)
    }
    
  }
}

func changePositionOfLine(x, y int, line *models.Line,g *Game) {
  line.StartX += float32(x)
  line.MaxX += float32(x)
  line.EndX = line.MaxX
  
  line.StartY += float32(y)
  line.MaxY += float32(y)
  line.EndY = line.MaxY

  px,py,ok:=checkColision(line, g)
  if ok{
    line.EndX=px
    line.EndY=py
  }

}

func changePositionOfLineRec(x, y int, line *models.Line,g *Game) {
  line.EndX = line.MaxX
  line.EndY = line.MaxY

  px,py,ok:= checkColision(line,g)
  if ok{
    line.EndX=px
    line.EndY= py
  }
}


func checkColision(line *models.Line, g *Game)(px,py float32 , ok bool){
  sides:=make([][2][2]float32,4)
  x0,y0:= g.rectangle.PositionX,g.rectangle.PositionY
  w,h := float32(g.rectangle.Width), float32(g.rectangle.Height)

  sides[0]=[2][2]float32{{x0, y0},{x0+w ,y0}}
  sides[1]=[2][2]float32{{x0 + w, y0},{x0+w ,y0 + h}}
  sides[2]=[2][2]float32{{x0 + w, y0 + h},{x0 ,y0 + h}}
  sides[3]=[2][2]float32{{x0, y0 + h },{x0, y0}}

  for _, v := range sides {
    x1, y1 := line.StartX, line.StartY
    x2, y2 := line.EndX, line.EndY
    x3, y3 := v[0][0], v[0][1]
    x4, y4 := v[1][0], v[1][1]

    den := (x1 - x2)*(y3 - y4) - (y1 - y2)*(x3 - x4)
    if den == 0 {
      continue
    }

    t := ((x1 - x3)*(y3 - y4) - (y1 - y3)*(x3 - x4)) / den
    u := ((x1 - x3)*(y1 - y2) - (y1 - y3)*(x1 - x2)) / den

    if t >= 0 && t <= 1 && u >= 0 && u <= 1 {

      px = x1 + t * (x2 - x1)
      py = y1 + t * (y2 - y1)
      line.EndX = px
      line.EndY = py
      return px,py,true
    }
  }
  
  return 0,0,false

}
