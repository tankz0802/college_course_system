package model

func GetDurationEquals64Rule(ac *AssignCourse) [][]*CourseSchedule {
	if len(ac.Time) == 0 {
		ac.Time = append(ac.Time, 1, 3)
	}
	if len(ac.Time) == 1 && !timeIncludesNight(ac.Time) {
		ac.Time = append(ac.Time, 3 - ac.Time[0])
	}else if len(ac.Time) == 1 && timeIncludesNight(ac.Time) {
		ac.Time = append(ac.Time, 2)
	}
	if len(ac.Time) == 3 {
		ac.Time = ac.Time[0:0]
		ac.Time = append(ac.Time, 1, 3)
	}
	res := getTimeCombine(ac.Time)

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
	res2 := getWeekDayCombine(res, ac.WeekDay)
	res3 := getWeekCombine(res2)
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

func getTimeCombine(time []int) [][][]int {
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
	var res [][][]int
	for _, item := range rule[0] {
		for _, item2 := range rule[1] {
			res = append(res, [][]int{item, item2})
		}
	}
	return res
}

func getWeekDayCombine(rule [][][]int, weekday []int) [][][]int {
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

func getWeekCombine(rule [][][]int) [][][]int {
	for _, item := range rule {
		itemCopy := make([][]int, 2)
		copy(itemCopy, item)
		if item[0][0] == 9 {
			item[0] = append(item[0], 1, 16)
			item[1] = append(item[1], 1, 8)
			itemCopy[0] = append(itemCopy[0], 1, 16)
			itemCopy[1] = append(itemCopy[1], 9, 16)
			rule = append(rule, itemCopy)
		}else if item[1][0] == 9 {
			item[0] = append(item[0], 1, 8)
			item[1] = append(item[1], 1, 16)
			itemCopy[0] = append(itemCopy[0], 9, 16)
			itemCopy[1] = append(itemCopy[1], 1, 16)
			rule = append(rule, itemCopy)
		}else{
			item[0] = append(item[0], 1, 16)
			item[1] = append(item[1], 1, 16)
		}
	}
	return rule
}

func timeIncludesNight(time []int) bool {
	for _, item := range time {
		if item == 3 {
			return true
		}
	}
	return false
}