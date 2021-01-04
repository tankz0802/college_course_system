package model

import "ccs/db"

type CourseSchedule struct {
	Id int64 `json:"id"`
	CId string `json:"cid"`
	WeekStart int32 `json:"week_start"`
	WeekEnd int32 `json:"week_end"`
	Day int32 `json:"day"`
	Start int32 `json:"start"`
	End int32 `json:"end"`
	TId int64 `json:"tid"`
}

func (cs *CourseSchedule) Add() error {
	sql := "insert into course_schedule(cid,week_start,week_end,day,start,end,tid) values(?,?,?,?,?,?,?,?)"
	_, err := db.DB.Exec(sql, cs.CId, cs.WeekStart, cs.WeekEnd, cs.Day, cs.Start, cs.End, cs.TId)
	if err != nil {
		return err
	}
	return nil
}

