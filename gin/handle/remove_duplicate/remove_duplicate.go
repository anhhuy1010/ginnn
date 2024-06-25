package removeduplicate

func RemoveDuplicate(nums []int) []int {
	a := []int{}
	m := make(map[int]int)
	for _, value := range nums {
		m[value] = value

	}
	for key := range m {
		a = append(a, key)
	}
	return a
}
