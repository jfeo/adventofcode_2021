package sonar

func DepthDifferences(measurements []int) (differences []int) {
	differences = make([]int, len(measurements)-1)
	for i := 1; i < len(measurements); i++ {
		differences[i-1] = measurements[i] - measurements[i-1]
	}
	return
}

func CountDepthIncreases(measurements []int) (increases int) {
	var differences = DepthDifferences(measurements)
	for i := 0; i < len(differences); i++ {
		if differences[i] > 0 {
			increases++
		}
	}
	return
}

func SlidingWindow(measurements []int, size int) (windows []int) {
	windows = make([]int, len(measurements)-size+1)
	for i := 0; i < len(measurements)-size+1; i++ {
		windows[i] = measurements[i]
		for j := 1; j < size; j++ {
			windows[i] += measurements[i+j]
		}
	}
	return
}
