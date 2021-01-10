package model

import (
	"ccs/db"
	"testing"
)

func TestClassCourse_Add(t *testing.T) {
	classCourse := &ClassCourse{
		Id: 1,
		Cid: "测试",
	}
	tx, err := db.DB.Begin()
	if err != nil {
		t.Error(err.Error())
	}
	err = classCourse.Add(tx)
	if err != nil {
		t.Error(err.Error())
	}
}