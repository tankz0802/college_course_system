package model

import (
	"ccs/db"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	CId      int64  `json:"cid"`
}

func (s *Student) UserIsExists() bool {
	sql := "select name from student where id=?"
	var name string
	_ = db.DB.QueryRow(sql, s.Id).Scan(&name)
	if name != "" {
		return true
	}
	return false
}

func (s *Student) Add() {
	sql := "insert into student values(?,?,?,?)"
	_, err := db.DB.Exec(sql, s.Id, s.Name, s.Password, s.CId)
	if err != nil {
		log.Println(err.Error())
	}
}

func (s *Student) CheckPassword() bool {
	sql := "select name,password from student where id=?"
	var password string
	_ = db.DB.QueryRow(sql, s.Id).Scan(&s.Name, &password)
	if s.Password == password {
		return true
	}
	return false
}

type StudentInfo struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
}

func GetAllStudentInfo() []*StudentInfo {
	sql := "select s.id,s.name,c.name from student as s inner join class as c on s.cid=c.id"
	rows, err := db.DB.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	studentInfoList := make([]*StudentInfo, 0)
	for rows.Next() {
		studentInfo := &StudentInfo{}
		if err := rows.Scan(studentInfo.Id, studentInfo.Name, studentInfo.Class); err != nil {
			log.Println(err)
			return nil
		}
		studentInfoList = append(studentInfoList, studentInfo)
	}
	return studentInfoList
}

type ClassStudent struct {
	Class       *Class     `json:"class"`
	StudentList []*Student `json:"student_list"`
}

func GetStudentList() ([]*ClassStudent, error) {

	sql := "select * from class"
	rows, err := db.DB.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	classList := make([]*Class, 0)
	for rows.Next() {
		class := &Class{}
		if err := rows.Scan(&class.Id, &class.Name, &class.TId); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		classList = append(classList, class)
	}
	classStudentList := make([]*ClassStudent, 0)
	for _, item := range classList {
		sql := "select id,name from student where cid=?"
		rows, err := db.DB.Query(sql, item.Id)
		if err != nil {
			rows.Close()
			log.Println(err.Error())
			return nil, err
		}
		classStudent := &ClassStudent{}
		classStudent.Class = item
		classStudent.StudentList = make([]*Student, 0)
		for rows.Next() {
			student := &Student{}
			if err := rows.Scan(&student.Id, &student.Name); err != nil {
				log.Println(err.Error())
				rows.Close()
				return nil, err
			}
			classStudent.StudentList = append(classStudent.StudentList, student)
		}
		rows.Close()
		classStudentList = append(classStudentList, classStudent)
	}
	return classStudentList, nil
}
