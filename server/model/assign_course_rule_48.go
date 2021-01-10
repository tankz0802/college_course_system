package model

func SelectDuration48Rule(ac *AssignCourse) int {
	if len(ac.Time) == 0 {
		ac.Time = append(ac.Time, 3)
	}
	if len(ac.Time) == 1 && !timeIncludesNight48(ac.Time) {
		ac.Time = append(ac.Time, 3 - ac.Time[0])
	}
	if timeIncludesNight48(ac.Time) {
		ac.Time = ac.Time[0:0]
		ac.Time = append(ac.Time, 3)
	}
	if len(ac.Time) == 3 {
		ac.Time = ac.Time[0:0]
		ac.Time = append(ac.Time, 3)
	}

	if len(ac.WeekDay) == 0 {
		ac.WeekDay = append(ac.WeekDay, 3, 5)
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
	if ac.Time[0] == 3 {
		return 1
	}else{
		return 2
	}
}

func GetDurationEquals48Rule1(ac *AssignCourse) []*CourseSchedule {
	courseScheduleList := make([]*CourseSchedule, 0)
	if ac.Time[0] == 3 {
		for _, item := range ac.WeekDay {
			courseSchedule := &CourseSchedule{
				CId: ac.Cid,
				WeekStart: 1,
				WeekEnd: 16,
				WeekDay: item,
				SectionStart: 9,
				SectionEnd: 11,
				TId: ac.Tid,
			}
			courseScheduleList = append(courseScheduleList, courseSchedule)
		}
	}
	return courseScheduleList
}

func GetDurationEquals48Rule2(ac *AssignCourse) [][]*CourseSchedule {

	res := getTimeCombine48(ac.Time)
	res2 := getWeekDayCombine48(res, ac.WeekDay)
	res3 := getWeekCombine48(res2)
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

func getTimeCombine48(time []int) [][][]int {
	var rule [][][]int
	for _, item := range time {
		if item == 1 {
			rule = append(rule, [][]int{{1, 2}, {3, 4}})
		}else if item == 2 {
			rule = append(rule, [][]int{{5, 6}, {7, 8}})
		}else{
			rule = append(rule, [][]int{{9, 11}})
		}
	}
	if len(rule) == 1 {
		return rule
	}
	var res [][][]int
	for _, item := range rule[0] {
		for _, item1 := range rule[1] {
			res = append(res, [][]int{item, item1})
		}
	}
	return res
}

func getWeekDayCombine48(rule [][][]int, weekday []int) [][][]int {
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

func getWeekCombine48(rule [][][]int) [][][]int {
	for _, item := range rule {
		itemCopy1 := make([][]int, 2)
		itemCopy2 := make([][]int, 2)
		itemCopy3 := make([][]int, 2)
		copy(itemCopy1, item)
		copy(itemCopy2, item)
		copy(itemCopy3, item)
		item[0] = append(item[0], 1, 16)
		item[1] = append(item[1], 1, 8)
		itemCopy1[0] = append(itemCopy1[0], 1, 16)
		itemCopy1[1] = append(itemCopy1[1], 9, 16)
		itemCopy2[0] = append(itemCopy2[0], 1, 8)
		itemCopy2[1] = append(itemCopy2[1], 1, 16)
		itemCopy3[0] = append(itemCopy3[0], 9, 16)
		itemCopy3[1] = append(itemCopy3[1], 1, 16)
		rule = append(rule, itemCopy1, itemCopy2, itemCopy3)
	}
	return rule
}

func timeIncludesNight48(time []int) bool {
	for _, item := range time {
		if item == 3 {
			return true
		}
	}
	return false
}