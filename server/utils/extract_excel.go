package utils

import (
	"ccs/model"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"path"
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
		sectionEndIndex := search(weekInfo, "节")
		section := weekInfo[1:sectionEndIndex]
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
	//fmt.Println("第", weekStart, "-", weekEnd, "周", Week[day], "第", start, "-", end, "节")
	return weekStart, weekEnd, day, start, end
}

func scheduleIsExists(courseScheduleList []*model.CourseSchedule, courseSchedule *model.CourseSchedule) bool {
	for _, schedule :=range courseScheduleList {
		if schedule.CId == courseSchedule.CId &&
			schedule.TId == courseSchedule.TId &&
			schedule.WeekStart == courseSchedule.WeekStart &&
			schedule.WeekEnd == courseSchedule.WeekEnd &&
			schedule.WeekDay == courseSchedule.WeekDay &&
			schedule.SectionStart == courseSchedule.SectionStart &&
			schedule.SectionEnd == courseSchedule.SectionEnd {
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

func courseGroupIsExists(courseGroupList []*model.CourseGroup, courseGroup *model.CourseGroup) bool {
	for _, item := range courseGroupList {
		if item.Id == courseGroup.Id {
			return true
		}
	}
	return false
}

func courseGroupTeacherIsExists(courseGroupTeacherList []*model.CourseGroupTeacher,
	courseGroupTeacher *model.CourseGroupTeacher) bool {
	for _, item := range courseGroupTeacherList {
		if item.CgId == courseGroupTeacher.CgId && item.TId == courseGroupTeacher.TId {
			return true
		}
	}
	return false
}

func extractExcel(path string, teacherCourseInfoList *TeacherCourseInfoList) error {
	xls, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	table := xls.GetRows("Sheet0")
	teacher := &model.Teacher{}
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
				teacher.Password = teacherId
				teacher.Title = 3
			}else if i >=2 && j >= 2 {
				course := &model.Course{}
				courseSchedule := &model.CourseSchedule{}
				teacherCourse := &model.TeacherCourse{}
				courseGroup := &model.CourseGroup{}
				courseGroupTeacher := &model.CourseGroupTeacher{}
				courseSchedule.TId = teacher.Id
				teacherCourse.TId = teacher.Id
				courseGroup.TId = teacher.Id
				courseGroupTeacher.TId = teacher.Id
				courseInfos := strings.Split(colTrim, "/")
				// 课程名称
				course.Name = courseInfos[0]
				course.StuNum = 0
				course.Status = 0
				course.MaxNum = 120
				courseGroup.Name = course.Name
				// 课程时间
				timeInfo := courseInfos[1]
				flag := 1
				if search([]rune(timeInfo), ")") == -1 && search([]rune(timeInfo), "周") == -1 {
					flag++
					timeInfo = courseInfos[flag]
					course.Name = courseInfos[0] + "/" + courseInfos[1]
					courseGroup.Name = course.Name
				}
				courseSchedule.WeekStart,
				courseSchedule.WeekEnd,
				courseSchedule.WeekDay,
				courseSchedule.SectionStart,
				courseSchedule.SectionEnd = getTimeStartAndEnd(courseInfos[flag], i, j)
				// 课程id
				idIndex := strings.Index(courseInfos[flag+2], ")-")
				if idIndex == -1 {
					return errors.New("找不到课程")
				}else{
					idIndex += 2
				}
				id := string([]byte(courseInfos[flag+2])[idIndex:])
				course.Id = id
				courseSchedule.CId = id
				teacherCourse.CId = id
				//课程组id
				cgId := string([]byte(id)[:9])
				courseGroup.Id, _ = strconv.ParseInt(cgId, 10, 64)
				courseGroupTeacher.CgId = courseGroup.Id
				// 课程类型
				course.Category = courseInfos[flag+4]
				if !courseIsExists(teacherCourseInfoList.CourseList, course) {
					teacherCourseInfoList.CourseList = append(teacherCourseInfoList.CourseList, course)
				}
				if !scheduleIsExists(teacherCourseInfoList.CourseScheduleList, courseSchedule) {
					teacherCourseInfoList.CourseScheduleList = append(teacherCourseInfoList.CourseScheduleList, courseSchedule)
				}
				if !teacherCourseIsExists(teacherCourseInfoList.TeacherCourseList, teacherCourse) {
					teacherCourseInfoList.TeacherCourseList = append(teacherCourseInfoList.TeacherCourseList, teacherCourse)
				}
				if !courseGroupIsExists(teacherCourseInfoList.CourseGroupList, courseGroup) {
					teacherCourseInfoList.CourseGroupList = append(teacherCourseInfoList.CourseGroupList, courseGroup)
				}
				if !courseGroupTeacherIsExists(teacherCourseInfoList.CourseGroupTeacherList, courseGroupTeacher) {
					teacherCourseInfoList.CourseGroupTeacherList = append(teacherCourseInfoList.CourseGroupTeacherList, courseGroupTeacher)
				}
			}
		}
	}
	teacherCourseInfoList.TeacherList = append(teacherCourseInfoList.TeacherList, teacher)
	return nil
}

type TeacherCourseInfoList struct {
	TeacherList []*model.Teacher `json:"teacher_list"`
	CourseList []*model.Course `json:"course_list"`
	CourseScheduleList []*model.CourseSchedule `json:"course_schedule_list"`
	TeacherCourseList []*model.TeacherCourse `json:"teacher_course_list"`
	CourseGroupList []*model.CourseGroup `json:"course_group_list"`
	CourseGroupTeacherList []*model.CourseGroupTeacher `json:"course_group_teacher"`
}

func ExtractExcels(filesPath string) *TeacherCourseInfoList {
	teacherCourseInfoList := &TeacherCourseInfoList{}
	teacherCourseInfoList.TeacherList = make([]*model.Teacher, 0)
	teacherCourseInfoList.CourseList = make([]*model.Course, 0)
	teacherCourseInfoList.CourseGroupList = make([]*model.CourseGroup, 0)
	teacherCourseInfoList.CourseScheduleList = make([]*model.CourseSchedule, 0)
	teacherCourseInfoList.CourseGroupTeacherList = make([]*model.CourseGroupTeacher, 0)
	filesList := listFiles(filesPath)
	for _, file :=range filesList {
		filePath := path.Join(filesPath, file)
		extractExcel(filePath, teacherCourseInfoList)
	}

	// 计算课时
	duration := make(map[string]int)
	for _, item := range teacherCourseInfoList.CourseScheduleList {
		if _, ok := duration[item.CId]; ok {
			duration[item.CId] += (item.WeekEnd - item.WeekStart + 1)*(item.SectionEnd-item.SectionStart+1)
		}else{
			duration[item.CId] = (item.WeekEnd - item.WeekStart + 1)*(item.SectionEnd-item.SectionStart+1)
		}
	}

	// 将学时和学分赋值到课程中
	for _, item := range teacherCourseInfoList.CourseList {
		item.Duration = duration[item.Id]
		item.Credit = float64(item.Duration)/16.0
	}
	return teacherCourseInfoList
}

type ClassInfo struct {
	Class []*model.Class
	StudentList [][]*model.Student
}

func ExtractStudents(path string, classs []string, tid int64) (*ClassInfo, error) {
	classInfo := &ClassInfo{
		Class: make([]*model.Class, 0),
		StudentList: make([][]*model.Student, 0),
	}

	for _, item := range classs {
		idRune := []rune(item)[2:]
		id, _ := strconv.ParseInt(string(idRune), 10, 64)
		class := &model.Class{
			Id: id,
			Name: item,
			TId: tid,
		}
		classInfo.Class = append(classInfo.Class, class)
		studentList, err := ExtractStudent(path, item, id)
		if err != nil {
			return nil, err
		}
		classInfo.StudentList = append(classInfo.StudentList, studentList)
	}
	return classInfo, nil
}

func ExtractStudent(path string, sheet string, cid int64) ([]*model.Student,error) {
	xls, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	table := xls.GetRows(sheet)
	studentList := make([]*model.Student, 0)
	for i, row := range table {
		if i < 3 {
			continue
		}else{
			id, _ := strconv.ParseInt(row[2], 10, 64)
			student := &model.Student{
				Id: id,
				Name: row[1],
				Password: row[2],
				CId: cid,
			}
			studentList = append(studentList, student)
		}
	}
	return studentList, nil
}

