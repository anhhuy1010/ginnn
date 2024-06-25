package square

func Squares(a []int) []int {

	new := make([]int, len(a))
	for i, val := range a {
		new[i] = val * val
	}
	n := len(new)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if new[j] > new[j+1] {
				new[j], new[j+1] = new[j+1], new[j]
			}
		}
	}
	return new
}
