// cmd/main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"tasker/internals/models"
	"tasker/internals/repository"
	"tasker/internals/service"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateTitle := updateCmd.String("title", "", "New title for the task")
	updateStatus := updateCmd.String("status", "", "New status for the task (e.g., pending, done)")

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
		var status string
		if len(os.Args) > 2 {
			status = os.Args[2]
		}
		tasks, err := service.ListTasks(models.TaskStatus(status))
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		for _, task := range tasks {
			fmt.Printf("[%d] %s - %s\n", task.ID, task.Title, task.Status)
		}

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: mark-done <id>")
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		msg, err := service.MarkStatus(uint(id), "done")

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(msg)

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: mark-in-progress <id>")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		msg, err := service.UpdateTask(uint(id), "", "in-progress")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(msg)
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

	case "update":
		// We expect: update <id> [--title <title>] [--status <status>]
		if len(os.Args) < 3 {
			fmt.Println("usage: update <id> [--title <title>] [--status <status>]")
			return
		}

		// Extract ID manually
		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("❌ invalid task ID:", idStr)
			return
		}

		// Parse flags AFTER the ID
		updateCmd.Parse(os.Args[3:])

		if *updateTitle == "" && *updateStatus == "" {
			fmt.Println("please provide at least one field to update (--title or --status)")
			return
		}

		var status models.TaskStatus
		if *updateStatus != "" {
			status = models.TaskStatus(*updateStatus)
		}

		msg, err := service.UpdateTask(uint(id), *updateTitle, status)
		if err != nil {
			fmt.Println("❌ error:", err)
		} else {
			fmt.Println("✅", msg)
		}

	default:
		fmt.Println("expected 'add' or 'list' subcommands")
	}
}
