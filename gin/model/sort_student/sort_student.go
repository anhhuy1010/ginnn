package sortstudent

import (
	"errors"
	requestStudent "gin/request/student"
)

func Sort(a []requestStudent.Student, sortby string, sortorder string) ([]requestStudent.Student, error) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sortby == "id" && sortorder == "up" {
				if a[j].Id > a[j+1].Id {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else if sortby == "name" && sortorder == "up" {
				if a[j].Name > a[j+1].Name {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else if sortby == "age" && sortorder == "up" {
				if a[j].Age > a[j+1].Age {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else if sortby == "class" && sortorder == "up" {
				if a[j].Class > a[j+1].Class {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else if sortby == "id" && sortorder == "down" {
				if a[j].Id < a[j+1].Id {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else if sortby == "name" && sortorder == "down" {
				if a[j].Name < a[j+1].Name {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else if sortby == "age" && sortorder == "down" {
				if a[j].Age < a[j+1].Age {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else if sortby == "class" && sortorder == "down" {
				if a[j].Class < a[j+1].Class {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else {
				return nil, errors.New("sortby or sortorder incorrect")
			}
		}
	}
	return a, nil
}
