package main

import (
	"fmt"
	"log"
	"os"
	//"strings"
	"todo/tasks"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("please try a command and text")
		return
	}
	
	
	//text := strings.Fields(os.Args[2])


	switch {
	case os.Args[1] == "add":
		err := tasks.AddTask(os.Args[2])
		if err != nil {
			log.Fatal("something wrong",err)
			return
		}
		fmt.Println("create task completed!!")

		}
		
	
	
}
