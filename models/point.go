package models

type Point [2]float32

func NewPoint(lat, lon float32) Point {
	return Point{lat, lon}
}

func (p Point) GetLon() float32 {
	return p[1]
}

func (p Point) GetLat() float32 {
	return p[0]
}
