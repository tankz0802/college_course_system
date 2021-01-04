package utils

import (
	"fmt"
	"io/ioutil"
)

func listFiles(pathname string) []string {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return nil
	}
	var files []string
	for _, fi := range rd {
		if fi.IsDir() {
			continue
		} else {
			fullName := fi.Name()
			files = append(files, fullName)
		}
	}
	return files
}

