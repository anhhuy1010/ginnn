package addstudent

import (
	"errors"
	requestStudent "gin/request/student"
)

func Add_Student(student []requestStudent.Student, person requestStudent.Student) ([]requestStudent.Student, error) {
	m := make(map[int]requestStudent.Student)
	for _, value := range student {
		m[value.Id] = value
	}
	_, ok := m[person.Id]
	if ok {

		return student, errors.New("ID already exists")

	} else {
		student = append(student, person)
	}

	return student, nil
}
