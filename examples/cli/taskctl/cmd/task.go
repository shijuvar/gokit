/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task <ID>",
	Short: "This command creates new task",
	Long:  `task code01 -t="create new repo" -d=01-feb-2022`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Task ID must be provided")
			return
		}
		category := chooseCategory()
		fmt.Println("Category:", category)
		fmt.Println("Task ID: ", args[0])
		//name, _ := cmd.Flags().GetString("name")
		title, _ := cmd.Flags().GetString("title")
		due, _ := cmd.Flags().GetString("due")
		fmt.Println("Title of the task :" + title)
		fmt.Println("Due date of the task :" + due)
	},
}

func init() {
	createCmd.AddCommand(taskCmd)
	taskCmd.PersistentFlags().StringP("title", "t", "", "Title of the task")
	taskCmd.PersistentFlags().StringP("due", "d", "", "Due Date")
	taskCmd.MarkPersistentFlagRequired("title")
}

// chooseCategory chooses category with promptui
func chooseCategory() string {
	items := []string{"Coding", "Learning", "Meeting", "Design", "R & D"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    "Select the category to add a task",
			Items:    items,
			AddLabel: "Add your own category",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}
