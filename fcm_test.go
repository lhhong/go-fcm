package fcm

import (
	"math"
	"testing"
)

type Element int

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

	interfaceSlice := make([]Interface, 3)
	interfaceSlice[0] = Element(2)
	interfaceSlice[1] = Element(3)
	interfaceSlice[2] = Element(5)
	Cluster(interfaceSlice, 3.6, 4)
}
