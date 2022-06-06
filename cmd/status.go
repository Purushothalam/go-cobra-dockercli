/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var Name string
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	//Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if Name != "" {
			fmt.Printf("Status of %s is Stopped", Name)
		} else if len(args) > 0 {
			argument := args[0]
			fmt.Printf("Status of %s is Running", argument)
		} else {
			fmt.Println("Argument (or) Flag is required to check the status")
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	statusCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "name to be displayed")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
