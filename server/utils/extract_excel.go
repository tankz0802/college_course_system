package utils

import (
	"ccs/model"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"strings"
)

var Week map[int]string
func init() {
	Week = make(map[int]string)
	Week[1] = "星期一"
	Week[2] = "星期二"
	Week[3] = "星期三"
	Week[4] = "星期四"
	Week[5] = "星期五"
	Week[6] = "星期六"
	Week[7] = "星期日"
}

func search(text []rune, what string) int {
	whatRunes := []rune(what)
	l1 := len(text)
	l2 := len(whatRunes)
	if l1 < l2 {
		return -1
	}
	i, j := 0, 0
	for ; i<=l1-l2; i++ {
		j = 0
		for ; j<l2; j++ {
			if text[i+j] != whatRunes[j] {
				break
			}
		}
		if j == l2 {
			return i
		}
	}
	return -1
}

func getTimeStartAndEnd(weekInfoStr string, row int, col int) (weekStart int, weekEnd int, day int, start int, end int) {
	weekInfo := []rune(weekInfoStr)
	weekStartIndex := search(weekInfo, ")")
	weekEndIndex := search(weekInfo, "周")
	if weekStartIndex != -1 {
		section := weekInfo[1:5]
		sectionIndex := strings.Index(string(section), "-")
		startStr := string(section[:sectionIndex])
		endStr := string(section[sectionIndex+1:])
		start, _ = strconv.Atoi(startStr)
		end, _ = strconv.Atoi(endStr)
	}else{
		start = (row-1)*2-1
		end = start + 1
	}
	day = col - 1
	week := weekInfo[weekStartIndex+1:weekEndIndex]
	weekIndex := strings.Index(string(week), "-")
	weekStartStr := string(week[:weekIndex])
	weekEndStr := string(week[weekIndex+1:])
	weekStart, _ = strconv.Atoi(weekStartStr)
	weekEnd, _ = strconv.Atoi(weekEndStr)
	fmt.Println("第", weekStart, "-", weekEnd, "周", Week[day], "第", start, "-", end, "节")
	return weekStart, weekEnd, day, start, end
}

func scheduleIsExists(courseScheduleList []*model.CourseSchedule, courseSchedule *model.CourseSchedule) bool {
	for _, schedule :=range courseScheduleList {
		if schedule.CId == courseSchedule.CId &&
			schedule.TId == courseSchedule.TId &&
			schedule.WeekStart == courseSchedule.WeekStart &&
			schedule.WeekEnd == courseSchedule.WeekEnd &&
			schedule.Day == courseSchedule.Day &&
			schedule.Start == courseSchedule.Start &&
			schedule.End == courseSchedule.End {
			return true
		}
	}
	return false
}

func courseIsExists(courseList []*model.Course, course *model.Course) bool {
	for _, item := range courseList {
		if item.Id == course.Id {
			return true
		}
	}
	return false
}

func teacherCourseIsExists(teacherCourseList []*model.TeacherCourse, teacherCourse *model.TeacherCourse) bool {
	for _, item := range teacherCourseList {
		if item.TId == teacherCourse.TId && item.CId == teacherCourse.CId {
			return true
		}
	}
	return false
}

func extractExcel(path string) *TeacherCourse {
	xls, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	table := xls.GetRows("Sheet0")
	tc := &TeacherCourse{}
	teacher := &model.Teacher{}
	courseList := make([]*model.Course, 0)
	courseScheduleList := make([]*model.CourseSchedule, 0)
	teacherCourseList := make([]*model.TeacherCourse, 0)
	for i, row := range table {
		for j, col := range row {
			colTrim := strings.ReplaceAll(col, " ", "")
			if colTrim == "" {
				continue
			}
			if i==0 && j==0 {
				colRune := []rune(colTrim)
				teacherNameStart := search(colRune, "学期")
				teacherNameEnd := search(colRune, "老师")
				teacherIdStart := search(colRune, "教工号:")
				teacherName := string(colRune[teacherNameStart+2: teacherNameEnd])
				teacherId := string(colRune[teacherIdStart+4:])
				teacher.Id, _  = strconv.ParseInt(teacherId, 10, 64)
				teacher.Name = teacherName
				teacher.Password = Sha256(teacherId)
				teacher.Title = 3
			}else if i >=2 && j >= 2 {
				course := &model.Course{}
				courseSchedule := &model.CourseSchedule{}
				teacherCourse := &model.TeacherCourse{}
				courseSchedule.TId = teacher.Id
				teacherCourse.TId = teacher.Id
				courseInfos := strings.Split(colTrim, "/")
				// 课程名称
				course.Name = courseInfos[0]
				// 课程时间
				courseSchedule.WeekStart,
				courseSchedule.WeekEnd,
				courseSchedule.Day,
				courseSchedule.Start,
				courseSchedule.End = getTimeStartAndEnd(courseInfos[1], i, j)
				// 课程id
				idIndex := strings.Index(courseInfos[3], "180")
				id := string([]byte(courseInfos[3])[idIndex:])
				course.Id = id
				courseSchedule.CId = id
				teacherCourse.CId = id
				// 课程类型
				course.Category = courseInfos[5]
				if !courseIsExists(courseList, course) {
					courseList = append(courseList, course)
				}
				if !scheduleIsExists(courseScheduleList, courseSchedule) {
					courseScheduleList = append(courseScheduleList, courseSchedule)
				}
				if !teacherCourseIsExists(teacherCourseList, teacherCourse) {
					teacherCourseList = append(teacherCourseList, teacherCourse)
				}
			}
		}
	}
	tc.Teacher = teacher
	tc.CourseList = courseList
	tc.CourseScheduleList = courseScheduleList
	tc.TeacherCourseList = teacherCourseList
	fmt.Println(tc.Teacher)
	for _, item := range tc.CourseList {
		fmt.Println(item)
	}
	for _, item := range tc.CourseScheduleList {
		fmt.Println(item)
	}
	for _, item := range tc.TeacherCourseList {
		fmt.Println(item)
	}
	return tc
}

type TeacherCourse struct {
	Teacher *model.Teacher
	CourseList []*model.Course `json:"course_list"`
	CourseScheduleList []*model.CourseSchedule `json:"course_schedule_list"`
	TeacherCourseList []*model.TeacherCourse `json:"teacher_course_list"`
}

func extractExcels(path string) []TeacherCourse {
	return nil
}
