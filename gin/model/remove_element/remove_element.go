package removeelement

func RemoveElement(nums []int, value int) (int, []int) {

	a := []int{}
	for _, val := range nums {
		if val != value {
			a = append(a, val)
		}
	}
	n := len(a)
	return n, a
}
