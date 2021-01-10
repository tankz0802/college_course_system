package model

import (
	"fmt"
	"testing"
)

func TestGetDurationEquals8Rule(t *testing.T) {
	ac := &AssignCourse{
		Cid: "12345",
		Tid: 1,
		Class: []int64{181, 182},
		Time: []int{1},
		WeekDay: []int{1},
	}
	res := GetDurationEquals8Rule(ac)
	for _, item := range res {
		fmt.Println(item)
	}
}

func TestSelectDuration16ule(t *testing.T) {
	ac := &AssignCourse{
		Cid: "12345",
		Tid: 1,
		Class: []int64{181, 182},
		Time: []int{1},
		WeekDay: []int{1},
	}
	res := GetDurationEquals16Rule(ac)
	for _, item := range res {
		fmt.Println(item)
	}
}

func TestGetDurationEquals32Rule2(t *testing.T) {
	ac := &AssignCourse{
		Cid: "12345",
		Tid: 1,
		Class: []int64{181, 182},
		Time: []int{1, 2},
		WeekDay: []int{1},
	}
	SelectDuration32Rule(ac)
	res := GetDurationEquals32Rule2(ac)
	for _, item := range res {
		fmt.Println(item[0], item[1])
	}
}

func TestGetDurationEquals32Rule1(t *testing.T) {
	ac := &AssignCourse{
		Cid: "12345",
		Tid: 1,
		Class: []int64{181, 182},
		Time: []int{1},
		WeekDay: []int{1},
	}
	fmt.Println(SelectDuration32Rule(ac))
	res := GetDurationEquals32Rule1(ac)
	for _, item := range res {
		fmt.Println(item)
	}
}

func TestGetDurationEquals48Rule1(t *testing.T) {
	ac := &AssignCourse{
		Cid: "12345",
		Tid: 1,
		Class: []int64{181, 182},
		Time: []int{3},
		WeekDay: []int{1},
	}
	SelectDuration48Rule(ac)
	res := GetDurationEquals48Rule1(ac)
	for _, item := range res {
		fmt.Println(item)
	}
}

func TestGetDurationEquals48Rule2(t *testing.T) {
	ac := &AssignCourse{
		Cid: "12345",
		Tid: 1,
		Class: []int64{181, 182},
		Time: []int{1},
		WeekDay: []int{1},
	}
	SelectDuration48Rule(ac)
	res := GetDurationEquals48Rule2(ac)
	for _, item := range res {
		fmt.Println(item[0], item[1])
	}
	//for _, item := range res {
	//	fmt.Println(item[0], item[1])
	//}
}

func TestGetDurationEquals64Rule(t *testing.T) {
	ac := &AssignCourse{
		Cid: "12345",
		Tid: 1,
		Class: []int64{181, 182},
		Time: []int{1, 3},
		WeekDay: []int{1},
	}
	res := GetDurationEquals64Rule(ac)
	for _, item := range res {
		fmt.Println(item[0], item[1])
	}
}

