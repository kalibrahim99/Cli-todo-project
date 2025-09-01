package tasks

import (
	"encoding/json"
	
	"log"
	"os"
)

type Task struct {
	ID int `json:"id"`
	Text string `json:"text"`
	Completed bool `json:"completed"`
}


func AddTask(text string) error {
	data , err := os.ReadFile("tasks.json")
  
	var currentTask []Task
	if err == nil {
		json.Unmarshal(data, &currentTask)
	} else {
		log.Fatal("smething wong :",err)
	}


	Newid := 1
	if len(currentTask) > 0 {
		lastask := currentTask[len(currentTask)-1]
		Newid = lastask.ID + 1
	}


	newTask := Task {
		ID: Newid,
		Text: text,
		Completed: false,
	}

	currentTask = append(currentTask, newTask)
	updateData, err1 := json.MarshalIndent(currentTask, "", " ")
	if err != nil {
		log.Fatal("try again",err1)
	}  

	return os.WriteFile("tasks.json",[]byte(updateData),0644)
}


