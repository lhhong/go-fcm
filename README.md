# go-fcm
Fuzzy C-Means Clustering for Golang allowing custom data types

## Usage
```go
import(
    "math"
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
    ...
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

func main() {

    // Map your own data slice to the structure implementing fcm.Interface
    fcmPoints := make([]fcm.Interface, len(points))
    for i, p := range points {
        fcmPoints[i] = FcmPoint(p)
    }

    // Retrieve centroids and weights of each (centroid, data point) pair
    centroids, weights := fcm.Cluster(fcmPoints, 2.0, 0.00001, 3)
}
```
