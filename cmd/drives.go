package cmd

import (
	"../internal"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(driveCmd)
}

var driveCmd = &cobra.Command{
	Use:   "drives",
	Short: "Print the available drives on the device",
	Long:  `Print the available drives on the device`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.AvailableFileSystems())
	},
}
