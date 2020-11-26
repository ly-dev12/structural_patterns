package main

import "fmt"

type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

type DrawShape struct{}

func (drawShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape")
}

type ICountour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}
type DrawCountour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

func (countour DrawCountour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Countour")
	countour.shape.drawShape(countour.x, countour.y)
}

func (countour DrawCountour) resizeByFactor(factor int) {
	countour.factor = factor
}

func main() {
	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}
	var countour ICountour = DrawCountour{x, y, DrawShape{}, 2}
	countour.drawContour(x, y)
}
