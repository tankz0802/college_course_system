package utils

import (
	"ccs/model"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"strings"
)

var week map[int]string
var section map[int]string
func init() {
	week = make(map[int]string)
	section = make(map[int]string)
	week[1] = "星期一"
	week[2] = "星期二"
	week[3] = "星期三"
	week[4] = "星期四"
	week[5] = "星期五"
	week[6] = "星期六"
	week[7] = "星期日"
	section[1] = "第一大节"
	section[2] = "第二大节"
	section[3] = "第三大节"
	section[4] = "第四大节"
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

func extractExcel(path string) *TeacherCourse {
	xls, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	table := xls.GetRows("Sheet0")
	tc := &TeacherCourse{}

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
				fmt.Println("任课老师:", teacherName)
				fmt.Println("教工号:", teacherId)
				tc.Teacher.Id, _  = strconv.ParseInt(teacherId, 10, 64)
				tc.Teacher.Name = teacherName
				tc.Teacher.Password = Sha256(teacherId)
				tc.Teacher.Title = 3
			}else if i >=2 && j >= 2 {
				courseInfos := strings.Split(colTrim, "/")
				// 课程名称
				fmt.Println("课程名称: ", courseInfos[0])
				// 课程时间
				courseTimeInfos := []rune(courseInfos[1])
				timeIndex := search(courseTimeInfos, ")")
				if timeIndex != -1 {
					fmt.Println("时间: " +
						"第",string(courseTimeInfos[timeIndex+1: timeIndex+5]), "周",
						week[j-1],
						"第", string(courseTimeInfos[1:5]), "节")
				}else{
					fmt.Println("时间: " +
						"第",string(courseTimeInfos[timeIndex+1: timeIndex+5]), "周",
						week[j-1],
						section[i-1])
				}
				// 课程id
				idIndex := strings.Index(courseInfos[3], "180")
				id := string([]byte(courseInfos[3])[idIndex:idIndex+12])
				fmt.Println("课程id: ", id)
				// 课程类型
				fmt.Println("课程类型: ", courseInfos[5])
			}
		}
	}
}

type TeacherCourse struct {
	Teacher *model.Teacher
	TeacherCourseList []*model.Course `json:"teacher_course_list"`
}

func extractExcels(path string) []TeacherCourse {

}
