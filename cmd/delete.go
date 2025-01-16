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

		if deleteAll {
			taskManager.Tasks = []util.Task{}
			taskManager.NextId = 1
			fmt.Println("All tasks deleted.")

			if err := taskManager.SaveToFile("temp.json"); err != nil {
				fmt.Println("Error saving tasks:", err)
			}
			return
		}

		if len(args) == 0 {
			fmt.Println("Error: You must provide task IDs to delete, or use the --all flag to delete all tasks.")
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

	deleteCmd.Flags().Bool("all", false, "Select all tasks for deletion")
}
