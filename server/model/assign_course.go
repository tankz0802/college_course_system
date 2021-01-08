package model

type AssignCourse struct {
	Cid string `json:"cid"`
	Tid int64 `json:"tid"`
	Class []int `json:"class"`
	Time []int `json:"time"`
	WeekDay []int `json:"weekday"`
}

func (ac *AssignCourse) Add() (bool, error) {
	course := &Course{Id: ac.Cid}
	err := course.Get()
	if err != nil {
		return false, err
	}
	GetTeacherCourseTable(ac.Tid)
	if err != nil {
		return false, err
	}
	classCourseTable := make([]*CourseSchedule, 0)
	for _, item := range ac.Class {
		courseSchedule, err := GetClassCourseTable(item)
		if err != nil {
			return false, err
		}
		classCourseTable = append(classCourseTable, courseSchedule...)
	}

	return true, nil
}

