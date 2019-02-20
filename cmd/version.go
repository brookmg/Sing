package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Sing",
	Long:  `All software has versions. This is Sing's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sing v0.0.1 -- HEAD")
	},
}