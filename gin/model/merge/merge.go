package merge

func Merge(arr1 []int, arr2 []int) []int {
	arr := []int{}
	arr = append(arr1, arr2...)
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
