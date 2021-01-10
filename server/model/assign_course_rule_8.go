package model

func GetDurationEquals8Rule(ac *AssignCourse) []*CourseSchedule {
	if len(ac.Time) == 0 || len(ac.Time) == 3 {
		ac.Time = append(ac.Time, 1)
	}
	if timeIncludesNight8(ac.Time) {
		ac.Time = ac.Time[0:0]
		ac.Time = append(ac.Time, 1)
	}
	if len(ac.WeekDay) == 0 {
		ac.WeekDay = append(ac.WeekDay, 2, 4)
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
	res := getTimeCombine8(ac.Time)
	res2 := getWeekDayCombine8(res, ac.WeekDay)
	res3 := getWeekCombine8(res2)
	courseScheduleList := make([]*CourseSchedule, 0)
	for _, item := range res3 {
		courseSchedule := &CourseSchedule{
			CId: ac.Cid,
			WeekStart: item[3],
			WeekEnd: item[4],
			WeekDay: item[2],
			SectionStart: item[0],
			SectionEnd: item[1],
			TId: ac.Tid,
		}
		courseScheduleList = append(courseScheduleList, courseSchedule)
	}
	return courseScheduleList
}

func getTimeCombine8(time []int) [][]int {
	var rule [][]int
	for _, item := range time {
		if item == 1 {
			rule = append(rule, []int{1, 2}, []int{3, 4})
		}else if item == 2 {
			rule = append(rule, []int{5, 6}, []int{7, 8})
		}
	}
	return rule
}

func getWeekDayCombine8(rule [][]int, weekday []int) [][]int {
	for i, item := range rule {
		var itemCopy = make([]int, 2)
		copy(itemCopy, item)
		rule[i] = append(rule[i], weekday[0])
		itemCopy = append(itemCopy, weekday[1])
		rule = append(rule, itemCopy)
	}
	return rule
}

func getWeekCombine8(rule [][]int) [][]int {
	for i, item := range rule {
		itemCopy1 := make([]int, 3)
		itemCopy2 := make([]int, 3)
		itemCopy3 := make([]int, 3)
		copy(itemCopy1, item)
		copy(itemCopy2, item)
		copy(itemCopy3, item)
		rule[i] = append(rule[i], 1, 4)
		itemCopy1 = append(itemCopy1, 5, 8)
		itemCopy2 = append(itemCopy2, 9, 12)
		itemCopy3 = append(itemCopy3, 13, 16)
		rule = append(rule, itemCopy1, itemCopy2, itemCopy3)
	}
	return rule
}

func timeIncludesNight8(time []int) bool {
	for _, item := range time {
		if item == 3 {
			return true
		}
	}
	return false
}