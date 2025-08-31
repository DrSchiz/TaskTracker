package functions

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Task struct {
	Id          int       "json:`id`"
	Description string    "json:`description`"
	Status      string    "json:`status`"
	CreatedAt   time.Time "json:`created_at`"
	UpdatedAt   time.Time "json:`updated_at`"
}

func Add(task string, file *os.File) {
	tasks := getTasks(file)
	var id int
	if len(tasks) == 0 {
		id = 1
	} else {
		id = getMaxId(tasks) + 1
	}

	newTask := Task{
		Id:          id,
		Description: task,
		Status:      "todo",
		CreatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	writeFile(tasks, file)
}

func ReadAll(file *os.File) {
	tasks := getTasks(file)

	for i, task := range tasks {
		fmt.Println("ID:", task.Id)
		fmt.Println("Описание:", task.Description)
		fmt.Println("Статус:", task.Status)
		fmt.Println("Создана:", task.CreatedAt)
		fmt.Println("Обновлена:", task.UpdatedAt)
		if (i + 1) != len(tasks) {
			fmt.Println("----------")
		}
	}
}

func Delete(id int, file *os.File) {
	tasks := getTasks(file)

	for i, task := range tasks {
		if id == task.Id {
			tasks = append(tasks[:i], tasks[i+1:]...)

			writeFile(tasks, file)
		}
	}
	fmt.Println("Нет такой задачи")

}

func Update(id int, newTask string, file *os.File) {
	tasks := getTasks(file)

	for i, task := range tasks {
		if id == task.Id {
			tasks[i].Description = newTask
			tasks[i].UpdatedAt = time.Now()
			writeFile(tasks, file)
		}
	}
}

func getTasks(file *os.File) []Task {
	path := string(file.Name())
	fileAbs, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
	}

	var fileData []byte
	fileData, err = os.ReadFile(fileAbs)
	if err != nil {
		fmt.Println(err)
	}

	var tasks []Task
	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		fmt.Println(err)
	}

	return tasks
}

func getMaxId(tasks []Task) int {
	var ids []int
	for _, task := range tasks {
		ids = append(ids, task.Id)
	}

	maxId := ids[0]
	for _, id := range ids {
		if maxId < id {
			maxId = id
		}
	}

	return maxId
}

func writeFile(tasks []Task, file *os.File) {
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println(err)
	}

	fileAbs, _ := filepath.Abs(file.Name())
	os.WriteFile(fileAbs, jsonData, 1)

	defer file.Close()

	os.Exit(1)
}
