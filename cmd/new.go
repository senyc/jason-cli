/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/senyc/jason-cli/pkg/jason"
	"github.com/senyc/jason-cli/pkg/tui"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
		and usage of using your command. For example:

		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(tui.InitialNewModel())
		pm, err := p.Run()
		if err != nil {
			return
		}
		pn := pm.(tui.NewModel)
		title := pn.Inputs[tui.Title].Value()
		date := pn.Inputs[tui.Date].Value()
		priority := pn.Inputs[tui.Priority].Value()
		priorityInt, _ := strconv.Atoi(priority)
		jason.AddNewTask(title, priorityInt, date)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
