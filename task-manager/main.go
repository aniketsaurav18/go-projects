package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      int8      `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const FILE_NAME string = "file.json"

func calmExit(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}

func handleCreateTask(s string) error {
	data, err := extract_data(FILE_NAME)
	if err != nil {
		panic(err)
	}
	var latestTaskId int = 0
	for _, v := range data {
		latestTaskId = max(latestTaskId, v.Id)
	}
	newTask := Task{
		Id:          latestTaskId + 1,
		Description: s,
		Status:      0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	data = append(data, newTask)
	err = updata_data(FILE_NAME, data)
	if err != nil {
		panic(err)
	}
	return nil
}

func handleUpdateTask(n int, s string) error {
	data, err := extract_data(FILE_NAME)
	if err != nil {
		panic(err)
	}
	if n > len(data) {
		calmExit("Please enter a valid task ID")
	}
	for i, d := range data {
		if d.Id == n {
			data[i].Description = s
			break
		}
	}
	err = updata_data(FILE_NAME, data)
	if err != nil {
		panic(err)
	}
	return nil
}

func handleDeleteTask(n int) error {
	data, err := extract_data(FILE_NAME)
	if err != nil {
		panic(err)
	}
	var newTaskData []Task
	deleted := false
	for i, d := range data {
		if d.Id == n {
			newTaskData = append(data[:i], data[i+1:]...)
			deleted = true
		}
	}
	if !deleted {
		calmExit("Task id does not exist.")
	}
	err = updata_data(FILE_NAME, newTaskData)
	if err != nil {
		panic(err)
	}
	return nil
}

func handleMarkInProgress(n int) error {
	data, err := extract_data(FILE_NAME)
	if err != nil {
		panic(err)
	}
	marked := false
	for i, d := range data {
		if d.Id == n {
			data[i].Status = 1
			marked = true
			break
		}
	}
	if !marked {
		calmExit("Task id does not exist")
	}
	err = updata_data(FILE_NAME, data)
	if err != nil {
		panic(err)
	}
	return nil
}

func handleMarkDone(n int) error {
	data, err := extract_data(FILE_NAME)
	if err != nil {
		panic(err)
	}
	marked := false
	for i, d := range data {
		if d.Id == n {
			data[i].Status = 2
			marked = true
			break
		}
	}
	if !marked {
		calmExit("Task id does not exist")
	}
	err = updata_data(FILE_NAME, data)
	if err != nil {
		panic(err)
	}
	return nil
}

func handleListTask(l string) error {
	data, err := extract_data(FILE_NAME)
	if err != nil {
		panic(err)
	}
	if l == "all" {
		for _, d := range data {
			status := "todo"
			if d.Status == 1 {
				status = "in-progress"
			} else if d.Status == 2 {
				status = "done"
			}
			fmt.Printf("Task-Id: %v, Desc: %v, Status: %v \n", d.Id, d.Description, status)
		}
		return nil
	}
	status := int8(0)
	if l == "done" {
		status = 2
	} else if l == "in-progress" {
		status = 1
	} else if l == "todo" {
		status = 0
	} else {
		calmExit("Unknown status")
	}
	for _, d := range data {
		if d.Status == status {
			fmt.Printf("Task-Id: %v, Desc: %v, Status: %v \n", d.Id, d.Description, l)
		}
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("please pass arguments")
		os.Exit(1)
	}
	switch args[1] {
	case "add":
		if len(args) > 3 {
			fmt.Println("Uses: add <task name>")
			return
		}
		fmt.Println(args[2])
		handleCreateTask(args[2])
	case "update":
		if len(args) > 4 {
			fmt.Println("Uses: update <task id> <updated task name>")
			return
		}
		i, err := strconv.Atoi(args[2])
		if err != nil {
			panic("Task id must be a number.")
		}
		err = handleUpdateTask(i, args[3])
		if err == nil {
			fmt.Printf("Successfully updated task id %v\n", i)
		}
	case "delete":
		if len(args) > 3 {
			calmExit("Uses: delete <task id>")
		}
		i, err := strconv.Atoi(args[2])
		if err != nil {
			calmExit("Task id must be a number")
		}
		handleDeleteTask(i)

	case "mark-in-progress":
		if len(args) > 3 {
			calmExit("Uses: mark-in-progress <task id>")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			calmExit("Task id must be a number")
		}
		handleMarkInProgress(id)

	case "mark-done":
		if len(args) > 3 {
			calmExit("Uses: mark-done <task id>")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			calmExit("Task id must be a number")
		}
		handleMarkDone(id)
	case "list":
		if len(args) > 3 {
			calmExit("Uses: list [todo | done | in-progress]")
		}
		if len(args) == 2 {
			handleListTask("all")
		} else {
			switch args[2] {
			case "todo":
				handleListTask("todo")
			case "done":
				handleListTask("done")
			case "in-progress":
				handleListTask("in-progress")
			}
		}
	default:
		calmExit("No such argument applicable")
	}
}
