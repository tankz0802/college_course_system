package model

import (
	"ccs/db"
	"log"
)

type Teacher struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Title int64 `json:"title"`
}

func (t *Teacher) TeacherIsExists() bool {
	sql := "select name from teacher where id=?"
	var name string
	_ = db.DB.QueryRow(sql, t.Id).Scan(&name)
	if name != "" {
		return true
	}
	return false
}

func (t *Teacher) Add() error {
	sql := "insert into teacher values(?,?,?,?)"
	_, err := db.DB.Exec(sql, t.Id, t.Name, t.Password, t.Title)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (t *Teacher) CheckPassword() bool {
	sql := "select name,password,title from teacher where id=?"
	var password string
	_ = db.DB.QueryRow(sql, t.Id).Scan(&t.Name, &password, &t.Title)
	if t.Password == password {
		return true
	}
	return false
}

func (t *Teacher) UpdateTitle(title int64) error {
	sql := "update teacher set title=? where id=?"
	_, err := db.DB.Exec(sql, title, t.Id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func GetAllTeacher() []*Teacher {
	sql := "select id,name from teacher"
	rows, err := db.DB.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return nil
	}
	teacherList := make([]*Teacher, 0)
	for rows.Next() {
		teacher := &Teacher{}
		if err := rows.Scan(teacher.Id, teacher.Name); err != nil {
			log.Println(err)
			return nil
		}
		teacherList = append(teacherList, teacher)
	}
	return teacherList
}