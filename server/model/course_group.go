package model

import (
	"ccs/db"
	"log"
)

type CourseGroup struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	TId int64 `json:"t_id"`
}

func (cg *CourseGroup) Add() error {
	sql := "insert into course_group values(?,?,?)"
	_, err := db.DB.Exec(sql, cg.Id, cg.Name, cg.TId)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
