package fcm

import (
	"log"
	"math"
	"testing"
)

type FcmFloat float64

func (e FcmFloat) Multiply(weight float64) Interface {
	return FcmFloat(float64(e) * weight)
}
func (e FcmFloat) Add(e2 Interface) Interface {
	return e + e2.(FcmFloat)
}
func (e FcmFloat) Norm(e2 Interface) float64 {
	return math.Abs(float64(e - e2.(FcmFloat)))
}

func TestFloat(t *testing.T) {

	interfaceSlice := make([]Interface, 6)
	interfaceSlice[0] = FcmFloat(2)
	interfaceSlice[1] = FcmFloat(3)
	interfaceSlice[2] = FcmFloat(4)
	interfaceSlice[3] = FcmFloat(5)
	interfaceSlice[4] = FcmFloat(15)
	interfaceSlice[5] = FcmFloat(16)
	centroid, weights := Cluster(interfaceSlice, 2.0, 0.00001, 2)
	for i, r := range weights {
		for j, w := range r {
			log.Printf("Ele %f, Cent %f, weight %f", interfaceSlice[j], centroid[i], w)
		}
	}
}

type Point struct {
	X float64
	Y float64
}

type FcmPoint Point

func (e FcmPoint) Multiply(weight float64) Interface {
	return FcmPoint{
		X: weight * e.X,
		Y: weight * e.Y,
	}
}

func (e FcmPoint) Add(e2 Interface) Interface {
	return FcmPoint{
		X: e2.(FcmPoint).X + e.X,
		Y: e2.(FcmPoint).Y + e.Y,
	}
}

func (e FcmPoint) Norm(e2 Interface) float64 {
	xDiff := e.X - e2.(FcmPoint).X
	yDiff := e.Y - e2.(FcmPoint).Y
	return math.Sqrt(math.Pow(xDiff, 2.0) + math.Pow(yDiff, 2.0))
}

func TestPoint(t *testing.T) {

	points := []Point{
		{X: 0.3, Y: 0.2},
		{X: 0.2, Y: 0.3},
		{X: 0.2, Y: 0.2},
		{X: 0.9, Y: 0.8},
		{X: 0.6, Y: 1.2},
		{X: 0.7, Y: 0.8},
		{X: 0.3, Y: 0.1},
		{X: 0.4, Y: 0.9},
		{X: 0.8, Y: 0.9},
		{X: 0.8, Y: 0.8},
		{X: 0.2, Y: 0.7},
		{X: 0.1, Y: 0.6},
		{X: 0.1, Y: 0.9},
		{X: 0.0, Y: 0.9},
		{X: 0.7, Y: 0.9},
		{X: 0.2, Y: 0.9},
		{X: 0.3, Y: 0.8},
	}

	fcmPoints := make([]Interface, len(points))
	for i, p := range points {
		fcmPoints[i] = FcmPoint(p)
	}
	centroid, weights := Cluster(fcmPoints, 2.0, 0.00001, 3)
	for i, r := range weights {
		for j, w := range r {
			log.Printf("Centroid (%f, %f), Element (%f, %f), weight %f",
				centroid[i].(FcmPoint).X, centroid[i].(FcmPoint).Y,
				fcmPoints[j].(FcmPoint).X, fcmPoints[j].(FcmPoint).Y, w)
		}
	}
}
