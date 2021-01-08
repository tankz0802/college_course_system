package utils

import (
	"fmt"
)

func InitData() {
	teacherCourseInfo := ExtractExcels("assets/teacher_course")
	for _, item := range teacherCourseInfo.TeacherList {
		item.Add()
	}
	fmt.Println("初始化教师数据成功!")

	for _, item := range teacherCourseInfo.CourseList {
		item.Add()
	}
	fmt.Println("初始化课程数据成功!")

	for _, item := range teacherCourseInfo.TeacherCourseList {
		item.Add()
	}
	fmt.Println("初始化老师-课程关联表成功!")

	for _, item := range teacherCourseInfo.CourseGroupList {
		item.Add()
	}
	fmt.Println("初始化课程组成功!")

	for _, item := range teacherCourseInfo.CourseGroupTeacherList {
		item.Add()
	}
	fmt.Println("初始化老师-课程组关联表成功!")

	//for _, item := range teacherCourseInfo.CourseScheduleList {
	//	fmt.Println(item)
	//	item.Add()
	//}
	//fmt.Println("初始化课程表数据成功!")

	classInfo, _ := ExtractStudents("assets/student.xlsx",
		[]string{"计科181", "计科182", "计科183", "计科184", "计科185", "计科186"},
		100103)

	for _, item := range classInfo.Class {
		item.Add()
	}
	fmt.Println("初始化班级数据成功")

	for _, class := range classInfo.StudentList {
		for _, student := range class {
			student.Add()
		}
	}
	fmt.Println("初始化学生数据成功")
}