package utils

import (
	"fmt"
	"testing"
)

func TestExtractStudents(t *testing.T) {
	res, _ := ExtractStudents("../assets/student.xlsx",
		[]string{"计科181", "计科182", "计科183", "计科184", "计科185", "计科186"},
		100103)

	for _, item := range res.Class {
		fmt.Println(item)
	}

	for _, class := range res.StudentList {
		for _, student := range class {
			fmt.Println(student)
		}
	}
}

func TestExtractStudent(t *testing.T) {
	res, _ := ExtractStudent("../assets/student.xlsx", "计科181", 181)
	for _, item := range res {
		fmt.Println(item)
	}
}

func TestExtractExcels(t *testing.T) {
	res := ExtractExcels("assets/teacher_course")
	for _, item := range res.TeacherList {
		//fmt.Println(item)
		item.Add()
	}
	fmt.Println("=======================================")
	for _, item := range res.CourseList {
		//fmt.Println(item)
		item.Add()
	}
	fmt.Println("=======================================")
	for _, item := range res.CourseScheduleList {
		fmt.Println(item)
		item.Add()
	}
	fmt.Println("=======================================")
	for _, item := range res.TeacherCourseList {
		//fmt.Println(item)
		item.Add()
	}
	fmt.Println("=======================================")
	for _, item := range res.CourseGroupList {
		//fmt.Println(item)
		item.Add()
	}
	fmt.Println("=======================================")
	for _, item := range res.CourseGroupTeacherList {
		//fmt.Println(item)
		item.Add()
	}
}

