package model

import (
	"ccs/db"
	"log"
)

type CourseSchedule struct {
	Id int64 `json:"id"`
	CId string `json:"cid"`
	WeekStart int `json:"week_start"`
	WeekEnd int `json:"week_end"`
	WeekDay int `json:"week_day"`
	SectionStart int `json:"section_start"`
	SectionEnd int `json:"section_end"`
	TId int64 `json:"tid"`
}

func (cs *CourseSchedule) Add() error {
	sql := "insert into course_schedule(cid,week_start,week_end,week_day,section_start,section_end,tid) values(?,?,?,?,?,?,?)"
	_, err := db.DB.Exec(sql, cs.CId, cs.WeekStart, cs.WeekEnd, cs.WeekDay, cs.SectionStart, cs.SectionEnd, cs.TId)
	if err != nil {
		return err
	}
	return nil
}


func CourseHasConflict(sid int64, cid string) (bool,error) {
	courseSchedule, err := GetCourseScheduleByCID(cid)
	if err != nil {
		return true, err
	}
	studentCourseTable, err := GetStudentCourseTable(sid)
	if err != nil {
		return true, err
	}
	for _, course := range studentCourseTable {
		for _,courseSchedule := range courseSchedule {
			if courseSchedule.WeekStart >= course.WeekStart && course.WeekStart <= course.WeekEnd &&
				courseSchedule.WeekDay == course.WeekDay &&
				courseSchedule.SectionStart >= course.SectionStart && courseSchedule.SectionStart <= course.SectionEnd{
				return true, nil
			}
		}
	}
	return false, nil
}

func GetClassCourseTable(cid int) ([]*CourseSchedule, error) {
	sql := `select cs.id,cs.week_start,cs.week_end,cs.week_day,cs.section_start,cs.section_end 
			from class_course as cc 
			inner join course_schedule as cs 
			on cc.cid=cs.cid where cc.id=?`
	rows, err := db.DB.Query(sql, cid)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	classCourseTable := make([]*CourseSchedule, 0)
	for rows.Next() {
		courseSchedule := &CourseSchedule{}
		if err := rows.Scan(&courseSchedule.Id,
			&courseSchedule.WeekStart,
			&courseSchedule.SectionEnd,
			&courseSchedule.WeekDay,
			&courseSchedule.SectionStart,
			&courseSchedule.SectionEnd);
		err != nil {
			log.Println(err.Error())
			return nil, err
		}
		classCourseTable = append(classCourseTable, courseSchedule)
	}
	return classCourseTable, nil
}

func GetCourseScheduleByCID(cid string) ([]*CourseSchedule, error) {
	sql := "select * from course_schedule where cid=?"
	rows, err := db.DB.Query(sql, cid)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	courseScheduleList := make([]*CourseSchedule, 0)
	for rows.Next() {
		courseSchedule := &CourseSchedule{}
		if err := rows.Scan(&courseSchedule.Id,
			&courseSchedule.CId,
			&courseSchedule.WeekStart,
			&courseSchedule.WeekEnd,
			&courseSchedule.WeekDay,
			&courseSchedule.SectionStart,
			&courseSchedule.SectionEnd,
			&courseSchedule.TId); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		courseScheduleList = append(courseScheduleList, courseSchedule)
	}
	return courseScheduleList, nil
}

type CourseScheduleInfo struct {
	Id int64 `json:"id"`
	CId string `json:"cid"`
	WeekStart int `json:"week_start"`
	WeekEnd int `json:"week_end"`
	WeekDay int `json:"week_day"`
	SectionStart int `json:"section_start"`
	SectionEnd int `json:"section_end"`
	Tname string `json:"tname"`
	Cname string `json:"cname"`
	Category string `json:"category"`
}

func GetTeacherCourseTable(tid int64) ([]*CourseScheduleInfo, error) {
	sql := `select cs.id,cs.cid,cs.week_start,cs.week_end,cs.week_day,cs.section_start,cs.section_end,t.name,c.name,c.category
			from course_schedule as cs
			inner join teacher as t on cs.tid=t.id 
			inner join course as c on cs.cid=c.id
			where tid=?`
	rows, err := db.DB.Query(sql, tid)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	courseScheduleInfoList := make([]*CourseScheduleInfo, 0)
	for rows.Next() {
		courseScheduleInfo := &CourseScheduleInfo{}
		if err := rows.Scan(&courseScheduleInfo.Id,
			&courseScheduleInfo.CId,
			&courseScheduleInfo.WeekStart,
			&courseScheduleInfo.WeekEnd,
			&courseScheduleInfo.WeekDay,
			&courseScheduleInfo.SectionStart,
			&courseScheduleInfo.SectionEnd,
			&courseScheduleInfo.Tname,
			&courseScheduleInfo.Cname,
			&courseScheduleInfo.Category);
			err != nil {
			log.Println(err.Error())
			return nil ,err
		}
		courseScheduleInfoList = append(courseScheduleInfoList, courseScheduleInfo)
	}
	return courseScheduleInfoList, nil
}

func GetStudentCourseTable(sid int64) ([]*CourseScheduleInfo,error) {
	sql := `select cs.id,cs.cid,cs.week_start,cs.week_end,cs.week_day,cs.section_start,cs.section_end,t.name,c.name,c.category
			from student_course as sc 
			inner join course_schedule as cs on sc.cid=cs.cid
			inner join teacher as t on cs.tid=t.id 
			inner join course as c on cs.cid=c.id
			where sid=?`
	rows, err := db.DB.Query(sql, sid)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	courseScheduleInfoList := make([]*CourseScheduleInfo, 0)
	for rows.Next() {
		courseScheduleInfo := &CourseScheduleInfo{}
		if err := rows.Scan(&courseScheduleInfo.Id,
			&courseScheduleInfo.CId,
			&courseScheduleInfo.WeekStart,
			&courseScheduleInfo.WeekEnd,
			&courseScheduleInfo.WeekDay,
			&courseScheduleInfo.SectionStart,
			&courseScheduleInfo.SectionEnd,
			&courseScheduleInfo.Tname,
			&courseScheduleInfo.Cname,
			&courseScheduleInfo.Category);
			err != nil {
			log.Println(err.Error())
			return nil ,err
		}
		courseScheduleInfoList = append(courseScheduleInfoList, courseScheduleInfo)
	}
	return courseScheduleInfoList, nil
}

func getCourseSchedule(cid string) ([]*CourseSchedule,error) {
	sql := "select * from course_schedule where cid=?"
	rows, err := db.DB.Query(sql, cid)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	courseScheduleList := make([]*CourseSchedule, 0)
	for rows.Next() {
		courseSchedule := &CourseSchedule{}
		if err := rows.Scan(&courseSchedule.Id,
			&courseSchedule.CId,
			&courseSchedule.WeekStart,
			&courseSchedule.WeekEnd,
			&courseSchedule.WeekDay,
			&courseSchedule.SectionStart,
			&courseSchedule.SectionEnd,
			&courseSchedule.TId);
		err != nil {
			log.Println(err.Error())
			return nil ,err
		}
		courseScheduleList = append(courseScheduleList, courseSchedule)
	}
	return courseScheduleList, nil
}

