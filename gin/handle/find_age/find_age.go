package findage

import (
	requestStudent "gin/request/student"
)

func FindAge(age int, b []requestStudent.Student) []requestStudent.Student {
	agee := []requestStudent.Student{}
	for _, value := range b {
		if age < value.Age {
			agee = append(agee, value)
		}
	}
	return agee
}
