package model

import "ccs/db"

type TeacherCourse struct {
	TId int64 `json:"tid"`
	CId string `json:"cid"`
}

func (tc *TeacherCourse) Add() error {
	sql := "insert into teacher_course values(?,?)"
	_, err := db.DB.Exec(sql, tc.TId, tc.CId)
	if err != nil {
		return err
	}
	return nil
}