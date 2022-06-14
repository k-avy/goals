package cmd

import (
	"fmt"

	"github.com/k-avy/goals/pkg/gols"
	"github.com/spf13/cobra"
)

func ListFolders() *cobra.Command {

	folderCmd := &cobra.Command{
		Use:   "folder",
		Short: "Command line tool to list folders",

		Run: func(cmd *cobra.Command, args []string) {

			path, _ := cmd.Flags().GetString("path")

			slice, err := gols.GetPathContent(path)
			if err != nil {
				fmt.Errorf("Wrong Path")
			}
			gols.PrintPathContent(slice, false)
		},
	}
	folderCmd.Flags().StringP("path", "p", ".", "list folder of particular path")
	return folderCmd
}
