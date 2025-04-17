package models

import "image/color"

type Line struct {
	StartX      float32
	StartY      float32
	EndX        float32
	EndY        float32
	StrokeWidth float32
	Color       color.Color
	Aa          bool
}
