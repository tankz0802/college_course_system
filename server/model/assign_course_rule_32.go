package model

import "fmt"

func SelectDuration32Rule(ac *AssignCourse) int {
	if len(ac.Time) == 0 {
		ac.Time = append(ac.Time, 1)
	}
	if timeIncludesNight32(ac.Time) {
		ac.Time = ac.Time[0:0]
		ac.Time = append(ac.Time, 1)
	}
	if len(ac.Time) == 3 {
		ac.Time = ac.Time[0:0]
		ac.Time = append(ac.Time, 2)
	}

	if len(ac.WeekDay) == 0 {
		ac.WeekDay = append(ac.WeekDay, 1, 3)
	}
	if len(ac.WeekDay) == 1 {
		if ac.WeekDay[0] > 3 {
			ac.WeekDay = append(ac.WeekDay, ac.WeekDay[0]-2)
		}else{
			ac.WeekDay = append(ac.WeekDay, ac.WeekDay[0]+2)
		}
	}
	if len(ac.WeekDay) > 2 {
		ac.WeekDay = ac.WeekDay[0:2]
	}
	if len(ac.Time) == 1 {
		return 1
	}else{
		return 2
	}
}

func GetDurationEquals32Rule1(ac *AssignCourse) []*CourseSchedule {
	courseScheduleList := make([]*CourseSchedule, 0)
	if ac.Time[0] == 1 {
		for _, item := range ac.WeekDay {
			courseSchedule1 := &CourseSchedule{
				CId: ac.Cid,
				WeekStart: 1,
				WeekEnd: 16,
				WeekDay: item,
				SectionStart: 1,
				SectionEnd: 2,
				TId: ac.Tid,
			}
			courseSchedule2 := &CourseSchedule{
				CId: ac.Cid,
				WeekStart: 1,
				WeekEnd: 16,
				WeekDay: item,
				SectionStart: 3,
				SectionEnd: 4,
				TId: ac.Tid,
			}
			courseScheduleList = append(courseScheduleList, courseSchedule1, courseSchedule2)
		}
	}else{
		for _, item := range ac.WeekDay {
			courseSchedule1 := &CourseSchedule{
				CId: ac.Cid,
				WeekStart: 1,
				WeekEnd: 16,
				WeekDay: item,
				SectionStart: 5,
				SectionEnd: 6,
				TId: ac.Tid,
			}
			courseSchedule2 := &CourseSchedule{
				CId: ac.Cid,
				WeekStart: 1,
				WeekEnd: 16,
				WeekDay: item,
				SectionStart: 7,
				SectionEnd: 8,
				TId: ac.Tid,
			}
			courseScheduleList = append(courseScheduleList, courseSchedule1, courseSchedule2)
		}
	}
	return courseScheduleList
}

func GetDurationEquals32Rule2(ac *AssignCourse) [][]*CourseSchedule {

	res := getTimeCombine32(ac.Time)
	res2 := getWeekDayCombine32(res, ac.WeekDay)
	fmt.Println(res2)
	res3 := getWeekCombine32(res2)
	courseScheduleLists := make([][]*CourseSchedule, 0)
	for _, item := range res3 {
		courseSchedule1 := &CourseSchedule{
			CId: ac.Cid,
			WeekStart: item[0][3],
			WeekEnd: item[0][4],
			SectionStart: item[0][0],
			SectionEnd: item[0][1],
			WeekDay: item[0][2],
			TId: ac.Tid,
		}
		courseSchedule2 := &CourseSchedule{
			CId: ac.Cid,
			WeekStart: item[1][3],
			WeekEnd: item[1][4],
			SectionStart: item[1][0],
			SectionEnd: item[1][1],
			WeekDay: item[1][2],
			TId: ac.Tid,
		}
		courseScheduleList := make([]*CourseSchedule, 0)
		courseScheduleList = append(courseScheduleList, courseSchedule1, courseSchedule2)
		courseScheduleLists = append(courseScheduleLists, courseScheduleList)
	}
	return courseScheduleLists
}

func getTimeCombine32(time []int) [][][]int {
	var rule [][][]int
	for _, item := range time {
		if item == 1 {
			rule = append(rule, [][]int{{1, 2}, {3, 4}})
		}else if item == 2 {
			rule = append(rule, [][]int{{5, 6}, {7, 8}})
		}
	}
	var res [][][]int
	for _, item := range rule[0] {
		for _, item1 := range rule[1] {
			res = append(res, [][]int{item, item1})
		}
	}
	return res
}

func getWeekDayCombine32(rule [][][]int, weekday []int) [][][]int {
	fmt.Println(rule)
	for _, item := range rule {
		var itemCopy = make([][]int, 2)
		copy(itemCopy, item)
		item[0] = append(item[0], weekday[0])
		item[1] = append(item[1], weekday[1])
		itemCopy[0] = append(itemCopy[0], weekday[1])
		itemCopy[1] = append(itemCopy[1], weekday[0])
		rule = append(rule, itemCopy)
	}
	return rule
}

func getWeekCombine32(rule [][][]int) [][][]int {
	for _, item := range rule {
		fmt.Println(item)
		itemCopy1 := make([][]int, 2)
		itemCopy2 := make([][]int, 2)
		itemCopy3 := make([][]int, 2)
		copy(itemCopy1, item)
		copy(itemCopy2, item)
		copy(itemCopy3, item)
		item[0] = append(item[0], 1, 8)
		item[1] = append(item[1], 1, 8)
		itemCopy1[0] = append(itemCopy1[0], 1, 8)
		itemCopy1[1] = append(itemCopy1[1], 9, 16)
		itemCopy2[0] = append(itemCopy2[0], 9, 16)
		itemCopy2[1] = append(itemCopy2[1], 1, 8)
		itemCopy3[0] = append(itemCopy3[0], 9, 16)
		itemCopy3[1] = append(itemCopy3[1], 9, 16)
		rule = append(rule, itemCopy1, itemCopy2, itemCopy3)
	}
	return rule
}

func timeIncludesNight32(time []int) bool {
	for _, item := range time {
		if item == 3 {
			return true
		}
	}
	return false
}