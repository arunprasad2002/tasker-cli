# Tasker

https://github.com/arunprasad2002/tasker-cli



Tasker is a simple command-line task manager written in Go. It stores tasks in a local JSON file and provides basic CRUD operations to help you track what needs doing.

## How to run

Clone the repository and run the following command:

```bash
git clone https://github.com/arunprasad2002/tasker-cli.git
cd tasker-cli
```

Run the following command to build and run the project:

```bash
go build -o tasker ./cmd
./tasker --help # To see the list of available commands

# To add a task
./tasker add "Buy groceries"

# To update a task
./tasker update 1 --title "Buy groceries and cook dinner"
./tasker update 1 --status completed
./tasker update 1 --title "Buy groceries and cook dinner" --status in_progress

# To delete a task
./tasker delete 1

# To mark a task as in progress/done
./tasker mark-in-progress 1
./tasker mark-done 1

# To list all tasks
./tasker list
./tasker list pending
./tasker list in_progress
./tasker list completed
```
