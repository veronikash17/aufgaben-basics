package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {
		if len(nums)==0{
		return 0
	}
	min := nums[0]
	for i:=1; i<len(nums); i++{
		if nums[i]<= min{
			min = nums[i]
		}
	}
	return min

}

func MinList2(nums []int) int {
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