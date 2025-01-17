package main

func maxPoints(points [][]int) int {
	pointCount := 0

	for i := 0; i < len(points); i++ {
		slopeMap := make(map[float64]int) // Reset map for each point
		duplicates := 0
		localMax := 0

		for j := i + 1; j < len(points); j++ {
			// Check for duplicate points
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				duplicates++
				continue
			}

			slope := calcSlope(points[i], points[j])
			slopeMap[slope]++
			if slopeMap[slope] > localMax {
				localMax = slopeMap[slope]
			}
		}
		// Update global max (including the current point and duplicates)
		pointCount = max(pointCount, localMax+duplicates+1)
	}

	return pointCount
}

// Corrected slope function
func calcSlope(p1, p2 []int) float64 {
	if p1[0] == p2[0] { // Handle vertical line case
		return float64(1e9) // Large number to represent infinity
	}
	return float64(p2[1]-p1[1]) / float64(p2[0]-p1[0])
}

// Max helper function
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
