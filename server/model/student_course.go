package model

import (
	"ccs/db"
	"log"
)

type StudentCourse struct {
	SId int64 `json:"sid"`
	CId string `json:"cid"`
	Grade int32 `json:"grade"`
	GPA float32 `json:"gpa"`
}

func (sc *StudentCourse) Add() error {
	sql := "insert into student_course values(?,?,?,?)"
	_, err := db.DB.Exec(sql, sc.SId, sc.CId, sc.Grade, sc.GPA)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}