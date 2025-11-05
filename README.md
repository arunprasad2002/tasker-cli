## Tasker

**Project URL:** https://github.com/arunprasad2002/tasker-cli

**Project Page:** https://arunprasad2002.github.io/tasker-cli

Tasker is a simple command‑line task manager written in Go. It stores tasks in a local JSON file and provides basic CRUD operations to help you track what needs doing.

### Features
- Add tasks
- List tasks (optionally filtered by status)
- Update title or status
- Mark tasks as done or in‑progress
- Delete tasks
- Persistent storage in `storage/tasks.json`

### Requirements
- Go (matching the version in `go.mod`)

### Project structure
```
cmd/
  main.go            # CLI entrypoint
internals/
  models/            # Domain models (Task, TaskStatus)
  repository/        # JSON file repository
  service/           # Business logic
storage/
  tasks.json         # Data store (auto‑created)
go.mod
```

### Build
```bash
go build -o tasker ./cmd
```

### Run without building
```bash
go run ./cmd --help
```

### Usage
Tasks are stored in `storage/tasks.json`. The repository ensures the folder/file exists when needed.

Statuses supported by the model are: `pending`, `in_progress`, `completed`.

#### Add a task
```bash
./tasker add "Buy milk"
# ✅ task added: Buy milk
```

#### List tasks
```bash
./tasker list
```

Filter by status (optional argument):
```bash
./tasker list pending
./tasker list in_progress
./tasker list completed
```

Output example:
```
[1] Buy milk - pending
[2] Write report - in_progress
```

#### Mark a task done
```bash
./tasker mark-done 1
```

#### Mark a task in‑progress
```bash
./tasker mark-in-progress 2
```

#### Update a task (title and/or status)
```bash
./tasker update <id> [--title <new title>] [--status <pending|in_progress|completed>]

# examples
./tasker update 3 --title "Write weekly report"
./tasker update 3 --status completed
./tasker update 3 --title "Write weekly report" --status in_progress
```

#### Delete a task
```bash
./tasker delete <id>
```

### Data file
- Location: `storage/tasks.json`
- Created automatically on first write
- You can back it up or version it if you want to persist tasks

### Development
- Run directly during development: `go run ./cmd <command> ...`
- Lint/test with your preferred tools; the code is plain Go with no external deps for storage.

### Notes
- Listing accepts an optional single status argument to filter results.
- The repository handles empty/nonexistent data files gracefully.

### License
Add your preferred license here.


