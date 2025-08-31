package main

import (
	"fmt"
	"os"
	"strconv"

	funcs "taskt/functions"
)

func main() {
	filepath := "tasks.json"
	if !checkFileExist(filepath) {
		os.Create(filepath)
	}

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}

	action := os.Args[1]
	switch action {
	case "add":
		task := os.Args[2]
		fmt.Println(os.Args)
		funcs.Add(task, file)
	case "ls":
		fmt.Println(os.Args)
		funcs.ReadAll(file)
	case "delete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
		}
		funcs.Delete(id, file)
	case "update":
		id, err := strconv.Atoi(os.Args[2])
		task := os.Args[3]
		if err != nil {
			fmt.Println(err)
		}
		funcs.Update(id, task, file)
	}

}

func checkFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	}
	return true
}
