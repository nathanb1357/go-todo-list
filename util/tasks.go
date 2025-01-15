package util

import (
	"encoding/json"
	"errors"
	"os"
)

// A task with a status, name, and ID
type Task struct {
	ID        int
	Name      string
	Completed bool
}

// A manager struct that stores a list of tasks
type TaskManager struct {
	Tasks  []Task
	NextId int
}

// Method that adds a new task to the list based on the description
func (manager *TaskManager) AddTask(name string) {
	if len(manager.Tasks) == 0 {
		manager.NextId = 1
	}
	newTask := Task{
		ID:        manager.NextId,
		Name:      name,
		Completed: false,
	}
	manager.Tasks = append(manager.Tasks, newTask)
	manager.NextId++
}

// Marks a task as complete
func (manager *TaskManager) CompleteTask(id int) error {
	for i, task := range manager.Tasks {
		if task.ID == id {
			manager.Tasks[i].Completed = true
			return nil
		}
	}
	return errors.New("task not found")
}

// Removes a task from the list
func (manager *TaskManager) RemoveTask(id int) error {
	for i, task := range manager.Tasks {
		if task.ID == id {
			manager.Tasks = append(manager.Tasks[:i], manager.Tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

// Saves the current task list into a specified file
func (manager *TaskManager) SaveToFile(fileName string) error {
	encoded, err := json.Marshal(manager)
	if err != nil {
		return err
	}
	return os.WriteFile("temp.json", encoded, 0644)
}

// Loads the current task list from a specified file
func (manager *TaskManager) LoadFromFile(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &manager)
}
