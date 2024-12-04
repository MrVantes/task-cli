# Task-CLI

A simple command-line tool built in Golang to track your tasks efficiently. https://roadmap.sh/projects/task-tracker

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

For Windows, here are the steps to build and install your Go application:

1. **Build the application**:
   If you have Go installed, you can build the application by opening a command prompt (or PowerShell) in the directory where your Go code is located and running:

   ```bash
   go build -o task-cli.exe
   ```

   This will generate the `task-cli.exe` executable.

2. **Move the executable to a directory in your PATH**:
   To make the application accessible system-wide, you should move it to a directory that's included in your `PATH`. Common directories for this purpose are `C:\Program Files` or `C:\Users\<your_username>\go\bin`, but you can choose another directory.

   You can manually move the `task-cli.exe` or use the command line:

   ```bash
   move task-cli.exe C:\path\to\desired\directory\
   ```

   For example:

   ```bash
   move task-cli.exe C:\Users\<your_username>\go\bin\
   ```

3. **Ensure the directory is in your PATH**:
   If the directory where you moved the executable is not already in your `PATH`, you can add it by following these steps:
   - Right-click on `This PC` or `Computer`, and select **Properties**.
   - Click on **Advanced system settings** on the left side.
   - In the **System Properties** window, click on the **Environment Variables** button.
   - Under **System variables**, scroll down and select the `Path` variable, then click **Edit**.
   - Click **New** and add the path to the directory where you moved `task-cli.exe` (e.g., `C:\Users\<your_username>\go\bin\`).
   - Click **OK** to save the changes.

After these steps, you should be able to run `task-cli` from any command prompt or PowerShell window.

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
