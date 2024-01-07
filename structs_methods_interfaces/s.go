package structsmethodsinterfaces

import "math"


type Rectangle struct {
	width float64
	height float64
}


type Circle struct {
	Radius float64	
}

type Triangle struct {
	Base float64
	Height float64
}


type Shape interface {
	Area() float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.height * rectangle.width
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius 
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}


func Perimeter(rectangle Rectangle) float64  {
	return 2 * (rectangle.width + rectangle.height)
}

