package cmd

import (
	"fmt"

	"github.com/k-avy/goals/pkg/gols"
	"github.com/spf13/cobra"
)

func ListCurrent() *cobra.Command {

	currentCmd := &cobra.Command{
		Use:   "current",
		Short: "Command line tool to print current path",

		Run: func(cmd *cobra.Command, args []string) {
			path := gols.FindCurrentPath()
			fmt.Println(path)
		},
	}

	return currentCmd
}
