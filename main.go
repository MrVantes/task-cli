package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const jsonFile = "tasks.json"

// Task struct represents a task.
type Task struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TaskData holds task arrays for unmarshaling from JSON.
type TaskData struct {
	ID          []int64     `json:"id"`
	Description []string    `json:"description"`
	Status      []string    `json:"status"`
	CreatedAt   []time.Time `json:"created_at"`
	UpdatedAt   []time.Time `json:"updated_at"`
}

func add(task string) {
	fileTask, err := os.ReadFile(jsonFile)
	if err != nil {
		if os.IsNotExist(err) {
			writeNewData(task)
			return
		}
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var data TaskData
	err = json.Unmarshal(fileTask, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	lastID := int64(0)
	if len(data.ID) > 0 {
		lastID = data.ID[len(data.ID)-1]
	}

	data.Description = append(data.Description, task)
	data.Status = append(data.Status, "not done")
	data.ID = append(data.ID, lastID+1)
	data.CreatedAt = append(data.CreatedAt, time.Now())
	data.UpdatedAt = append(data.UpdatedAt, time.Now())

	writeToFile(data)
	fmt.Printf("Task added (ID: %d)\n", lastID+1)
}

func list() {
	fileTask, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var data TaskData
	err = json.Unmarshal(fileTask, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	if len(data.ID) != len(data.Description) || len(data.ID) != len(data.Status) {
		fmt.Println("Data mismatch.")
		os.Exit(1)
	}

	var tasks []Task
	for i := 0; i < len(data.ID); i++ {
		tasks = append(tasks, Task{
			ID:          int64(i + 1),
			Description: data.Description[i],
			Status:      data.Status[i],
			CreatedAt:   data.CreatedAt[i],
			UpdatedAt:   data.UpdatedAt[i],
		})
	}

	var filteredTasks []Task
	if len(os.Args) < 3 || os.Args[2] == "" {
		filteredTasks = tasks
	} else {
		switch os.Args[2] {
		case "done":
			for _, task := range tasks {
				if task.Status == "done" {
					filteredTasks = append(filteredTasks, task)
				}
			}
		case "todo":
			for _, task := range tasks {
				if task.Status == "not done" {
					filteredTasks = append(filteredTasks, task)
				}
			}
		case "in-progress":
			for _, task := range tasks {
				if task.Status == "in-progress" {
					filteredTasks = append(filteredTasks, task)
				}
			}
		default:
			fmt.Println("Invalid filter. Use: done, todo, or in-progress.")
			os.Exit(1)
		}
	}

	for _, task := range filteredTasks {
		status := "✘"
		if task.Status == "done" {
			status = "✔"
		} else if task.Status == "in-progress" {
			status = "▶"
		}

		fmt.Printf(" %d) %-30s [%s]\n", task.ID, task.Description, status)
	}
}

func update(index int, task string) {
	fileTask, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var data TaskData
	err = json.Unmarshal(fileTask, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	if index < 0 || index >= len(data.ID) {
		fmt.Println("Invalid index:", index)
		return
	}

	data.Description[index-1] = task
	data.Status[index-1] = "not done"
	data.UpdatedAt[index-1] = time.Now()

	writeToFile(data)
	fmt.Printf("Task updated (Index: %d)\n", index)
}

func updateStatus(index int) {
	fileTask, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var data TaskData
	err = json.Unmarshal(fileTask, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	if index < 0 || index >= len(data.ID) {
		fmt.Println("Invalid index:", index)
		return
	}

	switch os.Args[1] {
	case "mark-done":
		data.Status[index-1] = "done"
	case "mark-in-progress":
		data.Status[index-1] = "in-progress"
	}
	data.UpdatedAt[index-1] = time.Now()

	writeToFile(data)
	fmt.Printf("Task updated (Index: %d)\n", index)
}

func delete(index int) {
	// Adjust for 1-based indexing if necessary
	zeroBasedIndex := index - 1

	// Read file data
	fileTask, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Unmarshal file data into a struct
	var data TaskData
	err = json.Unmarshal(fileTask, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	// Check if the adjusted index is valid (0-based index)
	if zeroBasedIndex < 0 || zeroBasedIndex >= len(data.ID) {
		fmt.Println("Invalid index:", index)
		return
	}

	// Perform deletion based on the adjusted index
	data.ID = append(data.ID[:zeroBasedIndex], data.ID[zeroBasedIndex+1:]...)
	data.Description = append(data.Description[:zeroBasedIndex], data.Description[zeroBasedIndex+1:]...)
	data.Status = append(data.Status[:zeroBasedIndex], data.Status[zeroBasedIndex+1:]...)
	data.CreatedAt = append(data.CreatedAt[:zeroBasedIndex], data.CreatedAt[zeroBasedIndex+1:]...)
	data.UpdatedAt = append(data.UpdatedAt[:zeroBasedIndex], data.UpdatedAt[zeroBasedIndex+1:]...)

	// Write the updated data back to the file
	writeToFile(data)

	// Print confirmation message
	fmt.Printf("Task deleted (Index: %d)\n", index)
}

func writeNewData(task string) {
	data := TaskData{
		ID:          []int64{1},
		Description: []string{task},
		Status:      []string{"not done"},
		CreatedAt:   []time.Time{time.Now()},
		UpdatedAt:   []time.Time{time.Now()},
	}
	writeToFile(data)
}

func writeToFile(data TaskData) {
	updatedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		os.Exit(1)
	}

	err = os.WriteFile(jsonFile, updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command.")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		task := strings.Join(os.Args[2:], " ")
		add(task)
	case "list":
		list()
	case "update":
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		task := strings.Join(os.Args[3:], " ")
		update(index, task)
	case "delete":
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		delete(index)
	case "mark-done":
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		updateStatus(index)
	case "mark-in-progress":
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		updateStatus(index)
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}
