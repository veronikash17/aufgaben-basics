package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}
