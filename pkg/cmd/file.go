package cmd

import (
	"fmt"

	"github.com/k-avy/goals/pkg/gols"
	"github.com/spf13/cobra"
)

func ListFiles() *cobra.Command {

	fileCmd := &cobra.Command{
		Use:   "file",
		Short: "Command line tool to list files",

		Run: func(cmd *cobra.Command, args []string) {

			path, _ := cmd.Flags().GetString("path")

			slice, err := gols.GetPathContent(path)
			if err != nil {
				fmt.Errorf("Wrong Path")
			}
			gols.PrintPathContent(slice, true)
		},
	}
	fileCmd.Flags().StringP("path", "p", ".", "list folder of particular path")
	return fileCmd
}
