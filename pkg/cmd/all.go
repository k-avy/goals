package cmd

import (
	"fmt"

	"github.com/k-avy/goals/pkg/gols"
	"github.com/spf13/cobra"
)

func ListAll() *cobra.Command {

	allCmd := &cobra.Command{
		Use:   "all",
		Short: "Command line tool to list files and folders",

		Run: func(cmd *cobra.Command, args []string) {
			path, _ := cmd.Flags().GetString("path")
			slice, err := gols.GetPathContent(path)
			if err != nil {
				fmt.Errorf("Wrong Path")
			}
			gols.PrintPathContent(slice, true)
			gols.PrintPathContent(slice, false)
		},
	}
	allCmd.Flags().StringP("path", "p", ".", "list folder of particular path")
	return allCmd
}
