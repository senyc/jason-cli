/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/senyc/jason-cli/pkg/auth"
	"github.com/senyc/jason-cli/pkg/tui"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(tui.InitiaAuthModel())
		pm, err := p.Run()
		if err != nil {
			return
		}
		pn := pm.(tui.AuthModel)
		apiKey := pn.Input.Value()
		err = auth.AddKeyToFS(apiKey)
		if err != nil {
			fmt.Print(err)
		}
		key, err := auth.GetKeyFromFS()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(key)
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
