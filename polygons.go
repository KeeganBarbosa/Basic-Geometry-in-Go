package main

import (
	"fmt"
)

type Points struct {
	X float64
	Y float64
}

type Polygon []Points

func dot(vec1, vec2 Points) float64 {
	return vec1.X*vec2.X + vec1.Y*vec2.Y
}

func cross(vec1, vec2 Points) float64 {
	return ((vec1.X * vec2.Y) - (vec1.Y * vec2.X))
}

func add(vec1, vec2 Points) Points {
	newvec := Points{
		X: vec1.X + vec2.X,
		Y: vec1.Y + vec2.Y,
	}
	return newvec
}

func diff(vec1, vec2 Points) Points {
	newvec := Points{
		X: vec1.X - vec2.X,
		Y: vec1.Y - vec2.Y,
	}
	return newvec
}

func main() {

	var point1 Points

	var point2 Points

	point1.X = 1

	point1.Y = 2

	point2.X = -3

	point2.Y = 5

	fmt.Println(add(point1, point2))
	fmt.Println(diff(point1, point2))
	fmt.Println(dot(point1, point2))
	fmt.Println(cross(point1, point2))

}
