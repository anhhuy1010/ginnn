package month

import "errors"

func GetDayInMonth(month int) ([]int, error) {
	var day int
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31

	case 4, 6, 9, 11:
		day = 30

	case 2:
		day = 28

	default:
		return nil, errors.New("month is invalid")

	}
	nums := []int{}
	for i := 0; i < day; i++ {
		nums = append(nums, i+1)
	}

	return nums, nil

}
