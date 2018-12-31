package fcm

import (
	"log"
	"math"
	"math/rand"
)

// Interface Define the operations within values
type Interface interface {
	Multiply(weight float64) Interface
	Add(v Interface) Interface
	Norm(v Interface) float64
}

// Cluster Perform FCM with randomly initialized centroids
func Cluster(vals []Interface, fuzziness float64, epsilon float64, numCentroids int) ([]Interface, [][]float64) {

	centroids := make([]Interface, 0, numCentroids)

	// TODO: Use map generation if ratio of numCentroids/len(val) is too small
	randIndices := rand.Perm(len(vals))[:numCentroids]

	for _, i := range randIndices {
		centroids = append(centroids, vals[i])
	}

	return centroids, ClusterGivenCentroids(vals, fuzziness, epsilon, centroids)
}

// ClusterGivenCentroids Perform FCM with manually initialized centroids
func ClusterGivenCentroids(vals []Interface, fuzziness float64, epsilon float64, centroids []Interface) [][]float64 {

	clusterWeights := make([][]float64, len(centroids))
	for i := 0; i < len(centroids); i++ {
		clusterWeights[i] = make([]float64, len(vals))
	}

	for evaluateWeights(vals, fuzziness, centroids, clusterWeights) > epsilon {
		recenter(vals, fuzziness, centroids, clusterWeights)
	}

	return clusterWeights
}

func evaluateWeights(vals []Interface, fuzziness float64, centroids []Interface, clusterWeights [][]float64) float64 {

	squareSum := 0.0
	for j, c := range centroids {
		for i, x := range vals {
			denominator := 0.0
			for _, c2 := range centroids {
				denominator += math.Pow((x.Norm(c) / x.Norm(c2)), 2.0/(fuzziness-1.0))
			}
			if math.IsNaN(denominator) {
				denominator = 1.0
			}
			newWeights := 1.0 / denominator
			squareSum += math.Pow((clusterWeights[j][i] - newWeights), 2.0)
			clusterWeights[j][i] = newWeights
		}
	}
	log.Printf("epsilon: %f", math.Sqrt(squareSum))
	return math.Sqrt(squareSum)
}

func recenter(vals []Interface, fuzziness float64, centroids []Interface, clusterWeights [][]float64) {

	for i := range centroids {
		normalization := 0.0
		var centroid Interface
		for j, w := range clusterWeights[i] {
			fuzziedWeight := math.Pow(w, fuzziness)

			normalization += fuzziedWeight

			adder := vals[j].Multiply(fuzziedWeight)
			if centroid == nil {
				centroid = adder
			} else {
				centroid = centroid.Add(adder)
			}

		}
		centroids[i] = centroid.Multiply(1.0 / normalization)

	}

}
