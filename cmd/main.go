// cmd/main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"tasker/internals/repository"
	"tasker/internals/service"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'list' subcommands")
		os.Exit(1)
	}

	repo := repository.NewTaskRepository("storage/tasks.json")
	service := service.NewTaskService(repo)

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if addCmd.NArg() < 1 {
			fmt.Println("please provide a task title")
			return
		}
		title := addCmd.Arg(0)
		if err := service.CreateTask(title); err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Println("✅ task added:", title)
		}

	case "list":
		listCmd.Parse(os.Args[2:])
		tasks, err := service.ListTasks()
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		for _, task := range tasks {
			fmt.Printf("[%d] %s - %s\n", task.ID, task.Title, task.Status)
		}

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if deleteCmd.NArg() < 1 {
			fmt.Println("please provide task ID to delete")
			return
		}

		idStr := deleteCmd.Arg(0)
		id, err := strconv.Atoi(idStr)

		if err != nil {
			fmt.Println("Invalid task ID:", idStr)
		}

		msg, err := service.DeleteTask(uint(id))

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(msg)
		}
	default:
		fmt.Println("expected 'add' or 'list' subcommands")
	}
}
