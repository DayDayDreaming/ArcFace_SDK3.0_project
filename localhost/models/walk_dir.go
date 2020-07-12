package models

import (
	"fmt"
	"io/ioutil"
)

func FileCount(path string) int {
	files,_ := ioutil.ReadDir(path)
	var count int
	for _, f := range files {
		fmt.Println(f.Name())
		count++
	}
	fmt.Println(count)
	return  count
}
