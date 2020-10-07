package course

/**
 * Max Pairwise Product Problem
 */


func maxPairwiseProductBrutal (arr []int) int {
	result := 0
	size := len(arr) - 1

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if (arr[i] * arr[j] > result) {
				result = arr[i] * arr[j]
			}
		}				
	}

	return result
}

func maxPairwiseProduct (arr []int) int {
	size := len(arr) - 1

	firstMaxIndex := -1
	for i := 0; i < size; i++ {
		if (firstMaxIndex == -1 || arr[i] > arr[firstMaxIndex]) {
			firstMaxIndex = i
		} 
	}

	secondMaxIndex := -1
	for i := 0; i < size; i++ {
		if (i != firstMaxIndex && (secondMaxIndex == -1 || arr[i] > arr[secondMaxIndex])) {
			secondMaxIndex = i
		} 
	}

	return arr[firstMaxIndex] * arr[secondMaxIndex]
}