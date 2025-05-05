package models

import "image/color"

type Rectangle struct{
  PositionX float32
  PositionY float32
  Height int
  Width int
  Color color.Color
  Aa bool
}
