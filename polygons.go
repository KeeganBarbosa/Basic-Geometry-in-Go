package main

import (
	"fmt"
)

type Points struct {
	X float64
	Y float64
}

type Polygon []Points

type Segment struct {
	beginning Points
	end       Points
}

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

// Slope is just diff, but makes context clearer from usecase.
// Slope from beginning to end

func slope(point1, point2 Points) Points {
	return diff(point2, point1)
}

func intersect(seg1, seg2 Segment) bool {
	eval := false
	s := slope(seg1.beginning, seg1.end)
	r := slope(seg2.beginning, seg2.end)
	t := cross((diff(seg1.beginning, seg2.beginning)), s) / (cross(r, s))
	if 0 <= t && t <= 1 {
		eval = !eval
		return eval
	}
	return eval
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
