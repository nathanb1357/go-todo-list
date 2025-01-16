/* Copyright © 2025 NATHAN BARTYUK */

package cmd

import (
	"fmt"
	"strconv"

	"github.com/nathanb1357/go-todo-list/util"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from your todo list",
	Run: func(cmd *cobra.Command, args []string) {
		deleteAll, _ := cmd.Flags().GetBool("all")
		deleteFinished, _ := cmd.Flags().GetBool("finished")

		if deleteAll {
			taskManager.Tasks = []util.Task{}
			taskManager.NextId = 1
			fmt.Println("All tasks deleted.")

			if err := taskManager.SaveToFile("temp.json"); err != nil {
				fmt.Println("Error saving tasks:", err)
			}
			return
		}

		if deleteFinished {
			var remainingTasks []util.Task
			for _, task := range taskManager.Tasks {
				if !task.Completed {
					remainingTasks = append(remainingTasks, task)
				}
			}
			taskManager.Tasks = remainingTasks
			fmt.Println("All finished tasks deleted.")

			if err := taskManager.SaveToFile("temp.json"); err != nil {
				fmt.Println("Error saving tasks:", err)
			}
			return
		}

		if len(args) == 0 {
			fmt.Printf("Error: You must provide task IDs, use the --all to delete all tasks, \nor use the --finished to delete all finished tasks.\n")
			return
		}

		var deletedIDs []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Error: %s is not a valid task ID.\n", arg)
				continue
			}

			err = taskManager.RemoveTask(id)
			if err != nil {
				fmt.Printf("Error: Task with ID %d not found.\n", id)
			} else {
				deletedIDs = append(deletedIDs, id)
			}
		}

		if len(deletedIDs) > 0 {
			fmt.Printf("Deleted tasks with IDs: %v\n", deletedIDs)

			// Save tasks to the JSON file
			if err := taskManager.SaveToFile("temp.json"); err != nil {
				fmt.Println("Error saving tasks:", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().BoolP("all", "a", false, "Delete all tasks")
	deleteCmd.Flags().BoolP("finished", "f", false, "Delete all finished tasks")
}
