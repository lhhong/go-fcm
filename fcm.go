package fcm

// Interface Define the operations within values
type Interface interface {
	Multiply(weight float64) Interface
	Add(v Interface) Interface
	Norm(v Interface) float64
}

// Cluster Perform FCM with randomly initialized centroids
func Cluster(vals []Interface, fuzziness float64, numCentroids int) ([]Interface, [][]float64) {

	centroids := make([]Interface, 0, numCentroids)

	return ClusterWithInitialCentroids(vals, fuzziness, centroids)
}

// ClusterWithInitialCentroids Perform FCM with manually initialized centroids
func ClusterWithInitialCentroids(vals []Interface, fuzziness float64, initCentroids []Interface) ([]Interface, [][]float64) {

	centroids := make([]Interface, 0, len(initCentroids))
	copy(centroids, initCentroids)
	clusterWeights := make([][]float64, len(centroids))
	for i := 0; i < len(centroids); i++ {
		clusterWeights = append(clusterWeights, make([]float64, len(vals)))
	}

	return centroids, clusterWeights
}
