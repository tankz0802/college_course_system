package model

import (
	"ccs/db"
	"log"
)

type Grade struct {
	Sid int64 `json:"sid"`
	Cid string `json:"cid"`
	Grade int64 `json:"grade"`
	GPA float64 `json:"gpa"`
}

func (g *Grade)Update() error {
	sql := "update student_course set grade=?,gpa=? where sid=? and cid=?"
	_, err := db.DB.Exec(sql, g.Grade, g.GPA, g.Sid, g.Cid)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

type GradeInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Teachers []string `json:"teachers"`
	Grade int `json:"grade"`
	GPA float32 `json:"gpa"`
}

func GetGradeInfoList(sid int64) ([]*GradeInfo,error) {
	sql := `select c.id,c.name,sc.grade,sc.gpa from student_course as sc 
			inner join course as c on sc.cid=c.id 
			where sc.sid=? 
			ORDER BY sc.grade desc;`
	rows, err := db.DB.Query(sql, sid)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	gradeInfoList := make([]*GradeInfo, 0)

	for rows.Next() {
		gradeInfo := &GradeInfo{}
		if err := rows.Scan(&gradeInfo.Id, &gradeInfo.Name, &gradeInfo.Grade, &gradeInfo.GPA); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		teachers, err := getTeacherNameByCid(gradeInfo.Id)
		if err != nil {
			return nil, err
		}
		gradeInfo.Teachers = teachers
		gradeInfoList = append(gradeInfoList, gradeInfo)
	}
	return gradeInfoList, nil
}

type StudentGrade struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Cname string `json:"cname"`
	Grade int `json:"grade"`
	GPA float32 `json:"gpa"`
}

func GetCourseStudentGradeList(cid string) ([]*StudentGrade,error) {
	sql := `select sc.sid,s.name,c.name,sc.grade,sc.gpa from student_course as sc 
			inner join student as s 
			on sc.sid=s.id 
			inner join class as c 
			on s.cid=c.id 
			where sc.cid=? 
			ORDER BY sc.grade desc,sc.sid asc`
	rows, err := db.DB.Query(sql, cid)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	courseStudentGradeList := make([]*StudentGrade, 0)
	for rows.Next() {
		studentGrade := &StudentGrade{}
		if err := rows.Scan(&studentGrade.Id,
			&studentGrade.Name,
			&studentGrade.Cname,
			&studentGrade.Grade,
			&studentGrade.GPA); err != nil {
			log.Println(err.Error())
			return nil ,err
		}
		courseStudentGradeList = append(courseStudentGradeList, studentGrade)
	}
	return courseStudentGradeList, nil
}




