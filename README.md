# Task-CLI

A simple command-line tool built in Golang to track your tasks efficiently.

## Installation

### Linux:

1. Build the application:
   ```bash
   go build -o task-cli
   ```
2. Move the executable to `/usr/local/bin/` for system-wide access:
   ```bash
   sudo mv task-cli /usr/local/bin/
   ```

### Windows:

- Download the precompiled binary from the release page or build it using the Go compiler.

---

## Usage

### Adding a New Task

To add a new task, use the following command:

```bash
task-cli add <task_description>
```

Example:

```bash
task-cli add "Buy groceries"
```

Output:

```
Task added successfully (ID: 1)
```

### Updating a Task

To update an existing task (using its ID), run:

```bash
task-cli update <task_id> <new_task_description>
```

Example:

```bash
task-cli update 1 "Buy groceries and cook dinner"
```

### Deleting a Task

To delete a task by ID:

```bash
task-cli delete <task_id>
```

Example:

```bash
task-cli delete 1
```

### Marking a Task as In Progress or Done

To mark a task as in progress:

```bash
task-cli mark-in-progress <task_id>
```

To mark a task as done:

```bash
task-cli mark-done <task_id>
```

### Listing Tasks

To list all tasks, simply use:

```bash
task-cli list
```

### Filtering Tasks by Status

You can filter tasks by status:

- **List all completed tasks**:

  ```bash
  task-cli list done
  ```

- **List all tasks yet to be started**:

  ```bash
  task-cli list todo
  ```

- **List tasks that are in progress**:
  ```bash
  task-cli list in-progress
  ```

---

## Contributing

Feel free to fork the repository, submit issues, or create pull requests. Contributions are welcome!

---

Let me know if you'd like to add or modify anything else!
