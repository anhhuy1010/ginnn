package removeid

import (
	"errors"
	requestStudent "gin/request/student"
)

func RemoveID(id int, b []requestStudent.Student) ([]requestStudent.Student, error) {
	removee := []requestStudent.Student{}
	r := false
	for _, value := range b {
		if id == value.Id {
			r = true
		} else {
			removee = append(removee, value)
		}
	}
	if !r {
		return removee, errors.New("id does not exist")
	}
	return removee, nil
}
