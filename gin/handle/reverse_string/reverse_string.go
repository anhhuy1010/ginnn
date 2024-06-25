package reversestring

func Reverse(arr string) string {
	var new string
	for i := len(arr) - 1; i >= 0; i-- {
		new = new + string(arr[i])
	}
	return new
}
