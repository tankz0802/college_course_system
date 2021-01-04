package model

import (
	"ccs/db"
	"log"
)

type Teacher struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Title int64 `json:"title"`
	Password string `json:"password"`
}

func (t *Teacher) TeacherIsExists() bool {
	sql := "select id from teacher where id=?"
	var id int64
	_ = db.DB.QueryRow(sql, t.Id).Scan(&id)
	if id > 0 {
		return true
	}
	return false
}

func (t *Teacher) Add() error {
	sql := "insert into teacher values(?,?,?,?)"
	_, err := db.DB.Exec(sql, t.Id, t.Name, t.Title, t.Password)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (t *Teacher) CheckPassword() bool {
	sql := "select password from teacher where id=?"
	var password string
	_ = db.DB.QueryRow(sql, t.Id).Scan(&password)
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