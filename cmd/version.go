package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short:"Print the version number of todo-cli",
	Long: "All software has versions. This is Todo's.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Todo cli application version: v1.0.")
	},
}
