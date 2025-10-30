// cmd/main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"tasker/internals/repository"
	"tasker/internals/service"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

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
	default:
		fmt.Println("expected 'add' or 'list' subcommands")
	}
}
