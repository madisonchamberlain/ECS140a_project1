package min
// Min returns the minimum value in the arr,
// and 0 if arr is nil.
func Min(arr []int) int {
	// test case nil in 0 out
	if arr == nil{
		return 0
	}
	// test case length = 0
	if len(arr) == 0{
		return 0
	}
	minVal := arr[0]
	// iterate through and find the minimum value from the list
	for _, curVal := range arr{
		if (curVal < minVal) {
			minVal = curVal
		}
	}
	// returns the minimum for the rest of the cases
	return minVal
}
