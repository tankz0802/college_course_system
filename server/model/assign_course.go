package model

import (
	"ccs/db"
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type AssignCourse struct {
	Cid string `json:"cid"`
	Tid int64 `json:"tid"`
	Class []int64 `json:"class"`
	Time []int `json:"time"`
	WeekDay []int `json:"weekday"`
}

func (ac *AssignCourse) Add() (bool, error) {
	course := &Course{Id: ac.Cid}
	err := course.Get()
	if course.Status == 1 {
		return false, errors.New("该课程已分配,无需重复操作")
	}
	if err != nil {
		return false, err
	}
	teacherCourseTable, err := GetTeacherCourseScheduleList(ac.Tid)
	if err != nil {
		return false, err
	}
	fmt.Println("老师课表")
	for _, item := range teacherCourseTable {
		fmt.Println(item)
	}
	classCourseTable := make([]*CourseSchedule, 0)
	for _, item := range ac.Class {
		courseSchedule, err := GetClassCourseTable(item)
		if err != nil {
			return false, err
		}
		classCourseTable = append(classCourseTable, courseSchedule...)
	}
	if len(ac.Class) != 0 {
		fmt.Println("班级课表")
		for _, item := range classCourseTable {
			fmt.Println(item)
		}
	}
	if course.Duration == 64 {
		assignCourseScheduleList := GetDurationEquals64Rule(ac)
		fmt.Println("按照偏好生成的排课规则:")
		for _, item := range assignCourseScheduleList {
			fmt.Println(item[0], item[1])
		}
		for _, item := range assignCourseScheduleList {
			if AssignCourseHasConflict(item[0], teacherCourseTable, classCourseTable) &&
				AssignCourseHasConflict(item[1], teacherCourseTable, classCourseTable) {
				continue
			}else{
				err := ac.AssignCourseToClass([]*CourseSchedule{item[0], item[1]})
				if err != nil {
					return false, err
				}
				return true, nil
			}
		}
		return false, errors.New("课程冲突较多,请选择新的规则")
	}else if course.Duration == 48 {
		if SelectDuration48Rule(ac) == 1 {
			assignCourseScheduleList := GetDurationEquals48Rule1(ac)
			fmt.Println("按照偏好生成的排课规则:")
			for _, item := range assignCourseScheduleList {
				fmt.Println(item)
			}
			for _, item := range assignCourseScheduleList {
				fmt.Println(item)
				if AssignCourseHasConflict(item, teacherCourseTable, classCourseTable) {
					continue
				}else{
					err := ac.AssignCourseToClass([]*CourseSchedule{item})
					if err != nil {
						return false, err
					}
					return true, nil
				}
			}
			return false, errors.New("课程冲突较多,请选择新的规则")
		}else{
			assignCourseScheduleList := GetDurationEquals48Rule2(ac)
			fmt.Println("按照偏好生成的排课规则:")
			for _, item := range assignCourseScheduleList {
				fmt.Println(item[0], item[1])
			}
			for _, item := range assignCourseScheduleList {
				fmt.Println(item[0], item[1])
				if AssignCourseHasConflict(item[0], teacherCourseTable, classCourseTable) &&
					AssignCourseHasConflict(item[1], teacherCourseTable, classCourseTable) {
					continue
				}else{
					err := ac.AssignCourseToClass([]*CourseSchedule{item[0], item[1]})
					if err != nil {
						return false, err
					}
					return true, nil
				}
			}
			return false, errors.New("课程冲突较多,请选择新的规则")
		}
	}else if course.Duration == 32 {
		if SelectDuration32Rule(ac) == 1 {
			assignCourseScheduleList := GetDurationEquals32Rule1(ac)
			fmt.Println("按照偏好生成的排课规则:")
			for _, item := range assignCourseScheduleList {
				fmt.Println(item)
			}
			for _, item := range assignCourseScheduleList {
				if AssignCourseHasConflict(item, teacherCourseTable, classCourseTable) {
					continue
				}
				err := ac.AssignCourseToClass([]*CourseSchedule{item})
				if err != nil {
					return false, err
				}
				return true, nil
			}
			return false, errors.New("课程冲突较多,请选择新的规则")
		}else{
			assignCourseScheduleList := GetDurationEquals32Rule2(ac)
			fmt.Println("按照偏好生成的排课规则:")
			for _, item := range assignCourseScheduleList {
				fmt.Println(item[0], item[1])
			}
			for _, item := range assignCourseScheduleList {
				if AssignCourseHasConflict(item[0], teacherCourseTable, classCourseTable) &&
					AssignCourseHasConflict(item[1], teacherCourseTable, classCourseTable) {
					continue
				}
				err := ac.AssignCourseToClass([]*CourseSchedule{item[0], item[1]})
				if err != nil {
					return false, err
				}
				return true, nil
			}
			return false, errors.New("课程冲突较多,请选择新的规则")
		}
	}else if course.Duration == 16 {
		assignCourseScheduleList := GetDurationEquals16Rule(ac)
		fmt.Println("按照偏好生成的排课规则:")
		for _, item := range assignCourseScheduleList {
			fmt.Println(item)
		}
		for _, item := range assignCourseScheduleList {
			if AssignCourseHasConflict(item, teacherCourseTable, classCourseTable) {
				continue
			}else{
				err := ac.AssignCourseToClass([]*CourseSchedule{item})
				if err != nil {
					return false, err
				}
				return true, nil
			}
		}
		return false, errors.New("课程冲突较多,请选择新的规则")
	}else if course.Duration == 8 {
		assignCourseScheduleList := GetDurationEquals8Rule(ac)
		fmt.Println("按照偏好生成的排课规则:")
		for _, item := range assignCourseScheduleList {
			fmt.Println(item)
		}
		for _, item := range assignCourseScheduleList {
			if AssignCourseHasConflict(item, teacherCourseTable, classCourseTable) {
				continue
			}else{
				err := ac.AssignCourseToClass([]*CourseSchedule{item})
				if err != nil {
					return false, err
				}
				return true, nil
			}
		}
		return false, errors.New("课程冲突较多,请选择新的规则")
	}
	return false, errors.New("课时不正确,请修改后重试")
}

func AssignCourseHasConflict(course *CourseSchedule, teacherCourse []*CourseSchedule, classCourse []*CourseSchedule) bool {
	for _,item := range teacherCourse {
		if course.WeekStart >= item.WeekStart && course.WeekStart <= item.WeekEnd &&
			course.WeekDay == item.WeekDay &&
			course.SectionStart >= item.SectionStart && course.SectionStart <= item.SectionEnd{
			return true
		}
	}
	for _,item := range classCourse {
		if course.WeekStart >= item.WeekStart && course.WeekStart <= item.WeekEnd &&
			course.WeekDay == item.WeekDay &&
			course.SectionStart >= item.SectionStart && course.SectionStart <= item.SectionEnd{
			return true
		}
	}
	return false
}

func (ac *AssignCourse) AssignCourseToClass(courseSchedule []*CourseSchedule) error {
	ctx, _:= context.WithTimeout(context.Background(), 10*time.Second)
	tx, err := db.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	sql := "update course set status=1 where id=?"
	_, err = tx.Exec(sql, ac.Cid)
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return err
	}
	fmt.Println("更新课程状态成功!")
	for _, item := range courseSchedule {
		err := item.Add(tx)
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}

	fmt.Println("插入课表成功!")
	for _, item := range ac.Class {
		classCourse := ClassCourse{
			Id: item,
			Cid: ac.Cid,
		}
		err := classCourse.Add(tx)
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}
	fmt.Println("插入班级课表成功!")
	studentCourses := make([]*StudentCourse, 0)
	for _, item := range ac.Class {
		sql := "select id from student where cid=?"
		rows, err := tx.Query(sql, item)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		for rows.Next() {
			var id int64
			if err := rows.Scan(&id); err != nil {
				log.Println(err.Error())
				return err
			}
			studentCourse := &StudentCourse{
				SId: id,
				CId: ac.Cid,
			}
			studentCourses = append(studentCourses, studentCourse)
		}
	}
	n := 0
	fmt.Println("获取学生名单成功!")
	for _, item := range studentCourses {
		//fmt.Println(item)
		err := item.Add(tx)
		n++
		if err != nil {
			tx.Rollback()
			log.Println(err.Error())
			return err
		}
	}
	fmt.Println(n)
	sql = "update course set stu_num=? where id=?"
	_, err = tx.Exec(sql, n, ac.Cid)
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return err
	}
	fmt.Println("更新课程人数成功")

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return err
	}
	fmt.Println("排课成功")

	return nil
}