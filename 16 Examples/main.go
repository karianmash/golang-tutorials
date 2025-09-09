package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const dataFile = "tasks.json"

func loadTasks() ([]Task, error) {
	var tasks []Task
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	tasks, _ := loadTasks()

	fmt.Println("Welcome to Todo CLI (type 'help' for commands, 'exit' to quit)")
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}
		if input == "exit" {
			fmt.Println("Goodbye 👋")
			break
		}

		args := strings.SplitN(input, " ", 2)
		command := args[0]

		switch command {
		case "help":
			fmt.Println("Commands: add <task>, list, done <id>, exit")

		case "add":
			if len(args) < 2 {
				fmt.Println("Please provide a task description")
				continue
			}
			title := args[1]
			task := Task{ID: len(tasks) + 1, Title: title, Done: false}
			tasks = append(tasks, task)
			err := saveTasks(tasks)
			if err != nil {
				return
			}
			fmt.Printf("Added task: %s\n", title)

		case "list":
			if len(tasks) == 0 {
				fmt.Println("No tasks found")
				continue
			}
			for _, t := range tasks {
				status := "❌"
				if t.Done {
					status = "✅"
				}
				fmt.Printf("[%d] %s %s\n", t.ID, status, t.Title)
			}

		case "done":
			if len(args) < 2 {
				fmt.Println("Please provide a task ID")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid task ID")
				continue
			}
			found := false
			for i, t := range tasks {
				if t.ID == id {
					tasks[i].Done = true
					err := saveTasks(tasks)
					if err != nil {
						return
					}
					fmt.Printf("Marked task %d as done\n", id)
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Task not found")
			}

		default:
			fmt.Println("Unknown command:", command)
		}
	}
}
