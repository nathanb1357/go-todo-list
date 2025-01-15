/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// finishCmd represents the finish command
var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		argument := args[0]
		id, err := strconv.Atoi(argument)
		if err != nil {
			fmt.Printf("Error: provided ID must be an integer!")
		} else {
			taskManager.CompleteTask(id)
			fmt.Printf("Task %d completed.\n", id)

			// Save tasks to the JSON file
			if err := taskManager.SaveToFile("temp.json"); err != nil {
				fmt.Println("Error saving tasks:", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(finishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// finishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// finishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
