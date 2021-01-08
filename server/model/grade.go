package model

import (
	"ccs/db"
	"log"
)

type Grade struct {
	Sid int64 `json:"sid"`
	Cid string `json:"cid"`
	Cname string `json:"cname"`
	Grade int64 `json:"grade"`
	GPA float64 `json:"gpa"`
}

func GetGradeListByStudent(sid int64) []*Grade {
	sql := "select sc.cid,sc.grade,sc.gpa,c.name from student_course as sc inner join course as c where sid=?"
	rows, err := db.DB.Query(sql, sid)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	gradeList := make([]*Grade, 0)
	for rows.Next() {
		grade := &Grade{}
		if err := rows.Scan(grade.Cid, grade.Grade, grade.GPA, grade.Cname); err != nil {
			log.Println(err)
			return nil
		}
		gradeList = append(gradeList, grade)
	}
	return gradeList
}

func GetGradeListByTeacher(cid string) []*Grade {
	sql := "select sc.sid,sc.cid,sc.grade,sc.gpa,c.name from student_course as sc inner join course as c where cid=?"
	rows, err := db.DB.Query(sql, cid)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	gradeList := make([]*Grade, 0)
	for rows.Next() {
		grade := &Grade{}
		if err := rows.Scan(grade.Sid, grade.Cid, grade.Grade, grade.GPA, grade.Cname); err != nil {
			log.Println(err)
			return nil
		}
		gradeList = append(gradeList, grade)
	}
	return gradeList
}

func UpdateGrade(sid int64, cid string, grade int64) error {
	sql := "update student_course set grade=? where sid=? and cid=?"
	_, err := db.DB.Exec(sql, grade, sid, cid)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}





