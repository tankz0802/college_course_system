package model

import (
	"fmt"
	"testing"
)

func TestGetDurationEquals64Rule(t *testing.T) {
	ac := &AssignCourse{
		Cid: "12345",
		Tid: 1,
		Class: []int{181, 182},
		Time: []int{1, 3},
		WeekDay: []int{1},
	}
	res := GetDurationEquals64Rule(ac)
	for _, item := range res {
		fmt.Println(item[0], item[1])
	}
}

