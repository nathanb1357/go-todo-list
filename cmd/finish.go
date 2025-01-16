/* Copyright © 2025 NATHAN BARTYUK */

package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// finishCmd represents the finish command
var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Mark a task on your todo list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		finishAll, _ := cmd.Flags().GetBool("all")

		if finishAll {
			for i := range taskManager.Tasks {
				taskManager.Tasks[i].Completed = true
			}
			fmt.Println("All tasks marked as compete.")

			if err := taskManager.SaveToFile("temp.json"); err != nil {
				fmt.Println("Error saving tasks:", err)
			}
			return
		}

		if len(args) == 0 {
			fmt.Println("Error: You must provide task IDs, or use the --all flag to finish all tasks.")
			return
		}

		var finishedIDs []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Error: %s is not a valid task ID.\n", arg)
				continue
			}

			err = taskManager.CompleteTask(id)
			if err != nil {
				fmt.Printf("Error: Task with ID %d not found.\n", id)
			} else {
				finishedIDs = append(finishedIDs, id)
			}
		}

		if len(finishedIDs) > 0 {
			fmt.Printf("Finished tasks with IDs: %v\n", finishedIDs)

			// Save tasks to the JSON file
			if err := taskManager.SaveToFile("temp.json"); err != nil {
				fmt.Println("Error saving tasks:", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(finishCmd)

	finishCmd.Flags().Bool("all", false, "Mark all tasks as complete")
}
