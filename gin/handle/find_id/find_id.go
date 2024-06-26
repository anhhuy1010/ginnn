package findid

import (
	"errors"
	requestStudent "gin/request/student"
)

func FindId(id int, b []requestStudent.Student) (requestStudent.Student, error) {
	for _, student := range b {
		if student.Id == id {
			return student, nil
		}
	}

	return requestStudent.Student{}, errors.New("ID does not exist")
}
