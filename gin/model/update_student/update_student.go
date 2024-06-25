package updatestudent

import (
	"errors"
	requestStudent "gin/request/student"
)

func Update(nums int, b []requestStudent.Student, p requestStudent.Student) ([]requestStudent.Student, error) {
	reg := false
	for i, value := range b {
		if value.Id == nums {
			b[i].Name = p.Name
			b[i].Age = p.Age
			b[i].Class = p.Class
			reg = true
			break
		}
	}
	if !reg {
		return nil, errors.New("id does not exist")
	}

	return b, nil
}
