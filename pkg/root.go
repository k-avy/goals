/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/k-avy/goals/pkg/cmd"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gols",
	Short: "Command line tool to list files and folders",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cmd.ListAll(), cmd.ListFiles(), cmd.ListFolders(), cmd.ListCurrent())
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
