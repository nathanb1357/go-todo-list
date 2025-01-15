/* Copyright © 2025 NATHAN BARTYUK */

package cmd

import (
	"fmt"
	"os"

	"github.com/nathanb1357/go-todo-list/util"
	"github.com/spf13/cobra"
)

var taskManager = util.TaskManager{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "go-todo-list",
	Version: "1.0.0",
	Short:   "A todo list to help keep track of many tasks.",
	Long: `A todo list CLI using Cobra to help keep track of tasks and 
offer persistent storage using JSON.`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Run: func(cmd *cobra.Command, args []string) {
		tasks := taskManager.Tasks
		fmt.Println("Tasks:")
		for _, task := range tasks {
			check := " "
			if task.Completed {
				check = "X"
			}
			fmt.Printf("%d\t\t[%s]\t\t%s\n", task.ID, check, task.Name)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	err := taskManager.LoadFromFile("temp.json")
	if err != nil {
		fmt.Println("No existing tasks found. Starting fresh.")
	}
}
