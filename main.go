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

type Task struct {
	ID          []int64     `json:"id"`
	Description []string    `json:"description"`
	Status      []string    `json:"status"`
	CreatedAt   []time.Time `json:"created_at"`
	UpdatedAt   []time.Time `json:"updated_at"`
}

func add(task string) {
	fileTask, err := os.ReadFile(jsonFile)
	if err != nil {
		// If the file doesn't exist, create a new Task object
		if os.IsNotExist(err) {
			fmt.Println("File not found, creating new data file...")
			writeNewData(task)
			return
		}
		fmt.Println("Error reading the file:", err)
		os.Exit(1)
	}

	// Parse the JSON data into a Data structure
	var data Task
	err = json.Unmarshal(fileTask, &data)
	if err != nil {
		fmt.Println("Error parsing the JSON:", err)
		os.Exit(1)
	}

	// Get the last ID and increment it
	lastID := int64(0)
	if len(data.ID) > 0 {
		lastID = data.ID[len(data.ID)-1] // Get the last ID from the array
	}

	// Add the new message to the Messages slice
	data.Description = append(data.Description, task)
	data.Status = append(data.Status, "not done")
	data.ID = append(data.ID, lastID+1)
	data.CreatedAt = append(data.CreatedAt, time.Now())
	data.UpdatedAt = append(data.UpdatedAt, time.Now())

	// Write the updated data back to the JSON file
	writeToFile(data)
	fmt.Printf("Task added successfully (ID: %d)\n ", data.ID[len(data.ID)-1])
}

func list() {
	// Read the tasks from the file
	fileTask, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		os.Exit(1)
	}

	// Parse the JSON data into a Data structure
	var data []Task
	err = json.Unmarshal(fileTask, &data)
	if err != nil {
		fmt.Println("Error parsing the JSON:", err)
		os.Exit(1)
	}

	// Declare a slice to hold the filtered tasks (or all tasks)
	var tasks []Task

	// If no argument is passed for the filter, show all tasks
	if len(os.Args) < 3 || os.Args[2] == "" {
		fmt.Println("No filter provided, showing all tasks.")
		tasks = data // If no filter is provided, just show all tasks
	} else {
		// Filtering based on command line arguments
		switch os.Args[2] {
		case "done":
			// Filter tasks with "done" status
			for _, task := range data {
				if contains(task.Status, "done") {
					tasks = append(tasks, task) // Reassign result of append
				}
			}
		case "todo":
			// Filter tasks with "not done" status
			for _, task := range data {
				if contains(task.Status, "not done") {
					tasks = append(tasks, task) // Reassign result of append
				}
			}
		case "in-progress":
			// Filter tasks with "in-progress" status
			for _, task := range data {
				if contains(task.Status, "in-progress") {
					tasks = append(tasks, task) // Reassign result of append
				}
			}
		default:
			fmt.Println("Invalid status filter. Please use: done, todo, or in-progress.")
			os.Exit(1)
		}
	}

	// Iterate through the filtered tasks and print them
	for _, task := range tasks {
		var status string
		// Set status based on the task's progress
		switch {
		case contains(task.Status, "done"):
			status = "✔"
		case contains(task.Status, "not done"):
			status = "✘"
		case contains(task.Status, "in-progress"):
			status = "▶"
		}

		// Format output with clear spacing and alignment
		fmt.Printf(" %d) %-30s [%s]\n", task.ID, task.Description, status)
	}
}

// Helper function to check if a slice contains a specific string
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func update(id int64, task string) {
	fileTask, err := os.ReadFile(jsonFile)
	if err != nil {
		// If the file doesn't exist, create a new Task object
		if os.IsNotExist(err) {
			fmt.Println("File not found, creating new data file...")
			writeNewData(task)
			return
		}
		fmt.Println("Error reading the file:", err)
		os.Exit(1)
	}

	// Parse the JSON data into a Data structure
	var data Task
	err = json.Unmarshal(fileTask, &data)
	if err != nil {
		fmt.Println("Error parsing the JSON:", err)
		os.Exit(1)
	}

	// Find the ID for updating
	updated := false
	for idx, taskID := range data.ID {
		if taskID == id {
			// Update the task at the found index
			data.Description[idx] = task
			data.Status[idx] = "not done"
			data.UpdatedAt[idx] = time.Now()

			// Mark that we found and updated the task
			updated = true
			fmt.Printf("Task updated successfully (ID: %d)\n ", id)

			// Break the loop since the task is updated
			break
		}
	}

	// If the task was not found, notify the user
	if !updated {
		fmt.Println("Task with ID", id, "not found.")
		return
	}

	// Write the updated data back to the file
	writeToFile(data)
}

func updateStatus(id int64) {
	fileTask, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		os.Exit(1)
	}

	// Parse the JSON data into a Data structure
	var data Task
	err = json.Unmarshal(fileTask, &data)
	if err != nil {
		fmt.Println("Error parsing the JSON:", err)
		os.Exit(1)
	}

	// Find the ID for updating
	updated := false
	for idx, taskID := range data.ID {
		if taskID == id {
			// Update the task at the found index
			switch os.Args[1] {
			case "mark-done":
				data.Status[idx] = "done"
				data.UpdatedAt[idx] = time.Now()
			case "mark-in-progress":
				data.Status[idx] = "in-progress"
				data.UpdatedAt[idx] = time.Now()
			default:
				data.Status[idx] = "not done"
			}

			// Mark that we found and updated the task
			updated = true
			fmt.Printf("Task updated successfully (ID: %d)\n ", id)

			// Break the loop since the task is updated
			break
		}
	}

	// If the task was not found, notify the user
	if !updated {
		fmt.Println("Task with ID", id, "not found.")
		return
	}

	// Write the updated data back to the file
	writeToFile(data)
}

func delete(id int64) {
	// Read json file
	fileTask, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		os.Exit(1)
	}

	// Parse the JSON data into a Data structure
	var data Task
	err = json.Unmarshal(fileTask, &data)
	if err != nil {
		fmt.Println("Error parsing the JSON:", err)
		os.Exit(1)
	}

	// Find the ID for deleting
	deleted := false
	for i, taskID := range data.ID {
		if taskID == id {
			// Remove the task by index
			data.ID = append(data.ID[:i], data.ID[i+1:]...)
			data.Description = append(data.Description[:i], data.Description[i+1:]...)
			data.Status = append(data.Status[:i], data.Status[i+1:]...)
			data.CreatedAt = append(data.CreatedAt[:i], data.CreatedAt[i+1:]...)
			data.UpdatedAt = append(data.UpdatedAt[:i], data.UpdatedAt[i+1:]...)

			// Mark that we found and preapre to delete the task
			deleted = true
			fmt.Printf("Task deleted successfully (ID: %d)\n ", id)

			// Break the loop since the task is delete
			break
		}
	}

	// If the task was not found, notify the user
	if !deleted {
		fmt.Println("Task with ID", id, "not found.")
		return
	}

	// Write the updated data back to the file
	writeToFile(data)
}

// Function to write new data to the file when the file doesn't exist
func writeNewData(task string) {
	data := Task{
		ID:          []int64{1},
		Description: []string{task},
		Status:      []string{"not done"},
		CreatedAt:   []time.Time{time.Now()},
		UpdatedAt:   []time.Time{time.Now()},
	}
	writeToFile(data)
}

// Function to write the updated data back to the JSON file
func writeToFile(data Task) {
	// Marshal the data to JSON
	updatedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		os.Exit(1)
	}

	// Write the updated data to the file
	err = os.WriteFile(jsonFile, updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}

}

func main() {
	// Check the number of arguments passed
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command.")
		os.Exit(1)
	}

	// Get the first argument as the command
	command := os.Args[1]

	// Use a switch statement to handle different commands
	switch command {
	case "add":
		// Set argument after 1 as strings
		task := strings.Join(os.Args[2:], " ")
		// Handle 'add' command
		add(task)
	case "list":
		// Handle 'list' command
		list()
	case "update":
		// Set ID for updating
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Println("Error converting id to int64:", err)
			os.Exit(1)
		}
		// Set argument after 2 as strings
		task := strings.Join(os.Args[3:], " ")
		// Handle 'update' command
		update(id, task)
	case "delete":
		// Set ID for updating
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Println("Error converting id to int64:", err)
			os.Exit(1)
		}
		// Handle 'delete' command
		delete(id)
	case "mark-done":
		// Set ID for updating
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Println("Error converting id to int64:", err)
			os.Exit(1)
		}
		// Handle 'update' command
		updateStatus(id)
	case "mark-in-progress":
		// Set ID for updating
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Println("Error converting id to int64:", err)
			os.Exit(1)
		}
		// Handle 'update' command
		updateStatus(id)
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}
