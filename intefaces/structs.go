package intefaces

import "math"
type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius  float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func RectanleArea(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
func CircleArea(circle Circle) float64 {
	return math.Pow(circle.Radius,2) * math.Pi 
}