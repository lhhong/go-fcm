package fcm_test

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/lhhong/go-fcm/fcm"
)

// Your current data structure
type Point struct {
	X float64
	Y float64
}

// Your current data
var points []Point = []Point{
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

// To implement Interface containing Multiply, Add and Norm
// Custom operators can be defined for different data types
type FcmPoint Point

// Multiplying a data point with a scalar weight
func (p FcmPoint) Multiply(weight float64) fcm.Interface {
	return FcmPoint{
		X: weight * p.X,
		Y: weight * p.Y,
	}
}

// Adding 2 data points together
func (p FcmPoint) Add(p2I fcm.Interface) fcm.Interface {
	p2 := p2I.(FcmPoint)
	return FcmPoint{
		X: p2.X + p.X,
		Y: p2.Y + p.Y,
	}
}

// Evaluating distance measure between 2 data points
func (p FcmPoint) Norm(p2I fcm.Interface) float64 {
	p2 := p2I.(FcmPoint)
	xDiff := p.X - p2.X
	yDiff := p.Y - p2.Y
	return math.Sqrt(math.Pow(xDiff, 2.0) + math.Pow(yDiff, 2.0))
}

func Example() {

	// Map your own data slice to the structure implementing fcm.Interface
	fcmPoints := make([]fcm.Interface, len(points))
	for i, p := range points {
		fcmPoints[i] = FcmPoint(p)
	}

	rand.Seed(10)

	// Retrieve centroids and weights of each (centroid, data point) pair
	centroids, weights := fcm.Cluster(fcmPoints, 2.0, 0.00001, 3)
	for i, r := range weights {
		for j, w := range r {
			fmt.Printf("Centroid (%f, %f), Element (%f, %f), weight %f\n",
				centroids[i].(FcmPoint).X, centroids[i].(FcmPoint).Y,
				fcmPoints[j].(FcmPoint).X, fcmPoints[j].(FcmPoint).Y, w)
		}
	}

	// Output: Centroid (0.182648, 0.831445), Element (0.300000, 0.200000), weight 0.006550
	// Centroid (0.182648, 0.831445), Element (0.200000, 0.300000), weight 0.035377
	// Centroid (0.182648, 0.831445), Element (0.200000, 0.200000), weight 0.006098
	// Centroid (0.182648, 0.831445), Element (0.900000, 0.800000), weight 0.046116
	// Centroid (0.182648, 0.831445), Element (0.600000, 1.200000), weight 0.276283
	// Centroid (0.182648, 0.831445), Element (0.700000, 0.800000), weight 0.029207
	// Centroid (0.182648, 0.831445), Element (0.300000, 0.100000), weight 0.025541
	// Centroid (0.182648, 0.831445), Element (0.400000, 0.900000), weight 0.661873
	// Centroid (0.182648, 0.831445), Element (0.800000, 0.900000), weight 0.007083
	// Centroid (0.182648, 0.831445), Element (0.800000, 0.800000), weight 0.017751
	// Centroid (0.182648, 0.831445), Element (0.200000, 0.700000), weight 0.889440
	// Centroid (0.182648, 0.831445), Element (0.100000, 0.600000), weight 0.682167
	// Centroid (0.182648, 0.831445), Element (0.100000, 0.900000), weight 0.952517
	// Centroid (0.182648, 0.831445), Element (0.000000, 0.900000), weight 0.879389
	// Centroid (0.182648, 0.831445), Element (0.700000, 0.900000), weight 0.014337
	// Centroid (0.182648, 0.831445), Element (0.200000, 0.900000), weight 0.974121
	// Centroid (0.182648, 0.831445), Element (0.300000, 0.800000), weight 0.899855
	// Centroid (0.248626, 0.209576), Element (0.300000, 0.200000), weight 0.989346
	// Centroid (0.248626, 0.209576), Element (0.200000, 0.300000), weight 0.948882
	// Centroid (0.248626, 0.209576), Element (0.200000, 0.200000), weight 0.990699
	// Centroid (0.248626, 0.209576), Element (0.900000, 0.800000), weight 0.030763
	// Centroid (0.248626, 0.209576), Element (0.600000, 1.200000), weight 0.077555
	// Centroid (0.248626, 0.209576), Element (0.700000, 0.800000), weight 0.014205
	// Centroid (0.248626, 0.209576), Element (0.300000, 0.100000), weight 0.956996
	// Centroid (0.248626, 0.209576), Element (0.400000, 0.900000), weight 0.068813
	// Centroid (0.248626, 0.209576), Element (0.800000, 0.900000), weight 0.003500
	// Centroid (0.248626, 0.209576), Element (0.800000, 0.800000), weight 0.010394
	// Centroid (0.248626, 0.209576), Element (0.200000, 0.700000), weight 0.064375
	// Centroid (0.248626, 0.209576), Element (0.100000, 0.600000), weight 0.236081
	// Centroid (0.248626, 0.209576), Element (0.100000, 0.900000), weight 0.022020
	// Centroid (0.248626, 0.209576), Element (0.000000, 0.900000), weight 0.062153
	// Centroid (0.248626, 0.209576), Element (0.700000, 0.900000), weight 0.005739
	// Centroid (0.248626, 0.209576), Element (0.200000, 0.900000), weight 0.010169
	// Centroid (0.248626, 0.209576), Element (0.300000, 0.800000), weight 0.037815
	// Centroid (0.756114, 0.871088), Element (0.300000, 0.200000), weight 0.004104
	// Centroid (0.756114, 0.871088), Element (0.200000, 0.300000), weight 0.015741
	// Centroid (0.756114, 0.871088), Element (0.200000, 0.200000), weight 0.003203
	// Centroid (0.756114, 0.871088), Element (0.900000, 0.800000), weight 0.923121
	// Centroid (0.756114, 0.871088), Element (0.600000, 1.200000), weight 0.646163
	// Centroid (0.756114, 0.871088), Element (0.700000, 0.800000), weight 0.956587
	// Centroid (0.756114, 0.871088), Element (0.300000, 0.100000), weight 0.017463
	// Centroid (0.756114, 0.871088), Element (0.400000, 0.900000), weight 0.269314
	// Centroid (0.756114, 0.871088), Element (0.800000, 0.900000), weight 0.989417
	// Centroid (0.756114, 0.871088), Element (0.800000, 0.800000), weight 0.971855
	// Centroid (0.756114, 0.871088), Element (0.200000, 0.700000), weight 0.046185
	// Centroid (0.756114, 0.871088), Element (0.100000, 0.600000), weight 0.081752
	// Centroid (0.756114, 0.871088), Element (0.100000, 0.900000), weight 0.025463
	// Centroid (0.756114, 0.871088), Element (0.000000, 0.900000), weight 0.058458
	// Centroid (0.756114, 0.871088), Element (0.700000, 0.900000), weight 0.979924
	// Centroid (0.756114, 0.871088), Element (0.200000, 0.900000), weight 0.015709
	// Centroid (0.756114, 0.871088), Element (0.300000, 0.800000), weight 0.062330
}
