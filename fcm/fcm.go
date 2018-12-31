// Package fcm performs Fuzzy C-Means clustering on custom data types
package fcm

import (
	"log"
	"math"
	"math/rand"
)

// Interface defines the data type that should be used with fcm. The data type is for a single data point and should implement the following:
// Multiply: Scale a data point to a given weight.
// Add: Add 2 data points together.
// Norm: Calculate distance measure between 2 data points.
type Interface interface {
	Multiply(weight float64) Interface
	Add(v Interface) Interface
	Norm(v Interface) float64
}

// Cluster performs FCM with randomly initialized centroids.
// vals is the slice of data points to be clustered.
// fuzziness is the hyperparameter (commonly known as m) determining level of fuzziness and should be > 1.
// epsilon determines the degree of convergence, where the algorithm terminates if change in cluster weights (measured by norm) < epsilon.
// numCentroids is the number of centroids to perform fcm with.
// It returns a slice of Interface and weight matrix where the first dimension refers to each centroid and the second dimension refering each data point,
// indicating weight of the data point belonging to the centroid.
func Cluster(vals []Interface, fuzziness float64, epsilon float64, numCentroids int) ([]Interface, [][]float64) {

	centroids := make([]Interface, numCentroids)

	initCentroids(vals, numCentroids, centroids)

	return centroids, ClusterGivenCentroids(vals, fuzziness, epsilon, centroids)
}

// ClusterGivenCentroids performs FCM with manually initialized centroids.
// Instead of numCentroids, the initial centroids will be passed in via centroids.
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

func initCentroids(vals []Interface, numCentroids int, centroids []Interface) {

	if len(vals) < 10000 || float64(numCentroids)/float64(len(vals)) > 0.2 {
		randIndices := rand.Perm(len(vals))[:numCentroids]

		for i, j := range randIndices {
			centroids[i] = vals[j]
		}
	} else {
		seen := map[int]bool{}

		for i := 0; i < numCentroids; i++ {
			randIndex := rand.Intn(len(vals))
			for _, present := seen[randIndex]; present; {
				randIndex = rand.Intn(len(vals))
			}
			seen[randIndex] = true
			centroids[i] = vals[randIndex]
		}
	}
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
