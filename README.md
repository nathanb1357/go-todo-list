# To-Do List CLI
A simple Command-Line Interface (CLI) application for managing tasks in a to-do list. This project is written in Go and leverages the Cobra library for command parsing and persistent task storage using JSON.

## Installation
1. **Clone the Repository:**
   ```
   git clone https://github.com/nathanb1357/go-todo-list.git
   cd go-todo-list
   ```
2. **Install Dependencies:** Ensure you have Go installed (version 1.19 or higher)
   ```
   go mod tidy
   ```
3. **Build the Application:**
   ```
   go build -o todo
   ```
4. **Run the Application:**
   ```
   ./todo
   ```
## JSON Storage
Tasks are stored in a JSON file (temp.json) in the project directory. The file is updated whenever you add, complete, or delete tasks. You can back up or modify this file as needed.

Example JSON structure:
```
[
  {
    "ID": 1,
    "Name": "Buy groceries",
    "Completed": false
  },
  {
    "ID": 2,
    "Name": "Finish project",
    "Completed": true
  }
]
```
