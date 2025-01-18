package main

func flippingMatrix(matrix [][]int32) int32 {
	n := len(matrix) / 2 // The size of the upper-left quadrant
	var sum int32 = 0

	// Iterate over the upper-left quadrant
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// Find the maximum value from the four candidates
			maxVal := max(
				matrix[i][j],
				matrix[2*n-1-i][j],
				matrix[i][2*n-1-j],
				matrix[2*n-1-i][2*n-1-j],
			)
			sum += maxVal
		}
	}

	return sum
}

// Helper function to find the maximum of four integers
func max(a, b, c, d int32) int32 {
	if a > b {
		b = a
	}
	if c > b {
		b = c
	}
	if d > b {
		b = d
	}
	return b
}
