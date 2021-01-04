package utils

import (
	"fmt"
	"testing"
)

func TestListFiles(t *testing.T) {
	path := "../assets/teacher_course"
	files := listFiles(path)
	fmt.Println(files)
}