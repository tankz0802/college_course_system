package model

import (
	"ccs/db"
	"log"
)

type Course struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Duration int `json:"duration"`
	Credit float64 `json:"credit"`
	Category string `json:"category"`
	StuNum int `json:"stu_num"`
	MaxNum int `json:"max_num"`
	Status int `json:"status"`
}

func (c *Course) CourseIsExists() bool {
	sql := "select name from course where id=?"
	var name string
	db.DB.QueryRow(sql, c.Id).Scan(&name)
	if name != "" {
		return true
	}
	return false
}

func (c *Course) Add() error {
	sql := "insert into course values(?,?,?,?,?,?,?,?)"
	_, err := db.DB.Exec(sql, c.Id, c.Name, c.Duration, c.Credit, c.Category, c.StuNum, c.MaxNum, c.Status)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (c *Course) Get() error {
	sql := "select * from course where id=?"
	row := db.DB.QueryRow(sql, c.Id)
	err := row.Scan(&c.Id, &c.Name, &c.Duration, &c.Credit, &c.Category, &c.StuNum, &c.MaxNum, &c.Status)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (c *Course) Update() error {
	sql := "update course set name=?,duration=?,credit=?,category=?,stu_num=?,max_num=? where id=?"
	_, err := db.DB.Exec(sql, c.Name, c.Duration, c.Credit, c.Category, c.StuNum, c.MaxNum, c.Id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func GetCourseList() ([]*Course, error) {
	sql := "select * from course"
	rows, err := db.DB.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	courseList := make([]*Course, 0)
	for rows.Next() {
		course := &Course{}
		if err := rows.Scan(&course.Id,
			&course.Name,
			&course.Duration,
			&course.Credit,
			&course.Category,
			&course.StuNum,
			&course.MaxNum,
			&course.Status);
		err != nil {
			log.Println(err.Error())
			return nil, err
		}
		courseList = append(courseList, course)
	}
	return courseList, nil
}

func GetUnTeachCourseList(tid int64) ([]*Course, error) {
	sql := `select tc.tid,c.id,c.name,c.duration,c.credit,c.category,c.stu_num,c.max_num
			from teacher_course as tc
			inner join course as c
			on tc.cid=c.id WHERE tc.tid=? and status=0`
	rows, err := db.DB.Query(sql, tid)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	teachCourseList := make([]*Course, 0)
	for rows.Next() {
		teachCourse := &Course{}
		var id int
		if err := rows.Scan(&id, &teachCourse.Id,
			&teachCourse.Name,
			&teachCourse.Duration,
			&teachCourse.Credit,
			&teachCourse.Category,
			&teachCourse.StuNum,
			&teachCourse.MaxNum);
			err != nil {
			return nil, err
		}
		teachCourseList = append(teachCourseList, teachCourse)
	}
	return teachCourseList, nil
}

func GetTeachCourseList(tid int64) ([]*Course, error) {
	sql := `select tc.tid,c.id,c.name,c.duration,c.credit,c.category,c.stu_num,c.max_num
			from teacher_course as tc
			inner join course as c
			on tc.cid=c.id WHERE tc.tid=? and status=1`
	rows, err := db.DB.Query(sql, tid)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	teachCourseList := make([]*Course, 0)
	for rows.Next() {
		teachCourse := &Course{}
		var id int
		if err := rows.Scan(&id, &teachCourse.Id,
			&teachCourse.Name,
			&teachCourse.Duration,
			&teachCourse.Credit,
			&teachCourse.Category,
			&teachCourse.StuNum,
			&teachCourse.MaxNum);
		err != nil {
			return nil, err
		}
		teachCourseList = append(teachCourseList, teachCourse)
	}
	return teachCourseList, nil
}

type CourseInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Duration int64 `json:"duration"`
	Credit float64 `json:"credit"`
	Category string `json:"category"`
	StuNum int64 `json:"stu_num"`
	MaxNum int64 `json:"max_num"`
	Teachers []string `json:"teachers"`
	Schedules []*CourseSchedule `json:"schedules"`
	Selected bool `json:"selected"`
}

func GetAllCourseInfo()([]*CourseInfo,error) {
	sql := "select * from course"
	rows, err := db.DB.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	courseInfoList := make([]*CourseInfo, 0)
	for rows.Next() {
		courseInfo := &CourseInfo{}
		if err := rows.Scan(&courseInfo.Id, &courseInfo.Name, &courseInfo.Duration, &courseInfo.Credit, &courseInfo.Category,
			&courseInfo.StuNum, &courseInfo.MaxNum); err != nil {
			log.Println(err)
			return nil, err
		}
		courseInfoList = append(courseInfoList, courseInfo)
	}

	for _, item := range courseInfoList {
		item.Teachers, err = getTeacherNameByCid(item.Id)
		if err != nil {
			return nil, err
		}
	}
	return courseInfoList,nil
}

func GetElectiveCourseInfo(sid int64) ([]*CourseInfo, error) {
	sql := `select c.id,c.name,c.duration,c.credit,c.category,c.stu_num,c.max_num,@selected:=IF(IFNULL(sc.sid,FALSE),TRUE, FALSE) as selected
			from course as c left join student_course as sc 
			on c.id=sc.cid 
			where c.category like '%选修%' and c.status=1 and (sc.sid=? or sc.cid is NULL)`
	rows, err := db.DB.Query(sql, sid)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	courseInfoList := make([]*CourseInfo, 0)
	for rows.Next() {
		courseInfo := &CourseInfo{}
		if err := rows.Scan(&courseInfo.Id, &courseInfo.Name, &courseInfo.Duration, &courseInfo.Credit, &courseInfo.Category,
			&courseInfo.StuNum, &courseInfo.MaxNum, &courseInfo.Selected); err != nil {
			log.Println(err)
			return nil, err
		}
		courseInfoList = append(courseInfoList, courseInfo)
	}
	for _, item := range courseInfoList {
		item.Teachers, err = getTeacherNameByCid(item.Id)
		if err != nil {
			return nil, err
		}

		item.Schedules, err = getCourseSchedule(item.Id)
		if err != nil {
			return nil, err
		}
	}
	return courseInfoList, nil
}

func getTeacherNameByCid(cid string) ([]string, error) {
	sql := "select t.name from teacher_course as tc inner join teacher as t on tc.tid=t.id where tc.cid=?"
	rows, err := db.DB.Query(sql, cid)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	teacherNameList := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		teacherNameList = append(teacherNameList, name)
	}
	return teacherNameList, nil
}

