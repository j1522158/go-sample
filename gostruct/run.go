package gostruct

import (
	"fmt"
	"unsafe"
)

type Task struct {
	Title    string
	Estimate int
}

func RunStruct() {

	fmt.Println("########## struct ##########")

	task1 := Task{
		Title:    "Learn Golang",
		Estimate: 3,
	}
	task1.Title = "Learning Go"

	fmt.Printf("task1-Title: %v task1-Estimate: %v\n", task1.Title, task1.Estimate)

	var task2 Task = task1
	task2.Title = "NEW"

	fmt.Printf("task1: %v task2: %v\n", task1.Title, task2.Title)

	task1p := &Task{
		Title:    "Learn concurrency",
		Estimate: 2,
	}
	fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))
	fmt.Printf("task1p: %T %v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))
}
