package model

import (
	"ccs/db"
	"log"
)

type CourseGroupTeacher struct {
	CgId int64 `json:"cg_id"`
	TId int64 `json:"tid"`
}

func (cgt *CourseGroupTeacher) Add() error {
	sql := "insert into course_group_teacher values(?,?)"
	_, err := db.DB.Exec(sql, cgt.CgId, cgt.TId)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}