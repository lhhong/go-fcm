package fcm

import (
	"log"
	"math"
	"math/rand"
	"testing"
)

type Element float64

func (e Element) Multiply(weight float64) Interface {
	return Element(float64(e) * weight)
}
func (e Element) Add(e2 Interface) Interface {
	return e + e2.(Element)
}
func (e Element) Norm(e2 Interface) float64 {
	return math.Abs(float64(e - e2.(Element)))
}

func TestBuild(t *testing.T) {

	rand.Seed(125)

	interfaceSlice := make([]Interface, 6)
	interfaceSlice[0] = Element(2)
	interfaceSlice[1] = Element(3)
	interfaceSlice[2] = Element(4)
	interfaceSlice[3] = Element(5)
	interfaceSlice[4] = Element(15)
	interfaceSlice[5] = Element(16)
	centroid, weights := Cluster(interfaceSlice, 2.0, 0.00001, 2)
	for i, r := range weights {
		for j, w := range r {
			log.Printf("Ele %f, Cent %f, weight %f", interfaceSlice[j], centroid[i], w)
		}
	}
}
