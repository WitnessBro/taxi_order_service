package models

type Point [2]float32

func NewPoint(lat, lon float32) Point {
	return Point{lat, lon}
}

func getLon(p Point) float32 {
	return p[1]
}

func getLat(p Point) float32 {
	return p[0]
}
