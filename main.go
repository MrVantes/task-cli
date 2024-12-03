package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const jsonFile = "tasks.json"

type Task struct {
	ID          []int64  `json:"id"`
	Description []string `json:"description"`
	Status      []string `json:"status"`
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

	// Write the updated data back to the JSON file
	writeToFile(data)
	fmt.Printf("Task added successfully (ID: %d)\n ", data.ID[len(data.ID)-1])
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
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}
