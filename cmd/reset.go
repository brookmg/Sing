package cmd

import (
	"../internal"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resetCmd)
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Command to reset music indexing",
	Long:  `Command to reset the indexing of music in the device. Basically refreshing the list`,
	//TraverseChildren:true,
	Run: func(cmd *cobra.Command, args []string) {
		internal.GetRefreshedMusicList(internal.AvailableFileSystems())
		fmt.Println("The music indexing has been refreshed.")
	},
}
