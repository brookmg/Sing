package cmd

import (
	"../internal"
	"github.com/spf13/cobra"
)

var port int

func init() {
	serverCmd.Flags().IntVarP(&port, "port" , "p", 15477, "set port of the server")
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts a server at port",
	Long:  `Starts a server on the local machine on a specified port`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.StartServer(internal.AvailableFileSystems(), port)
	},
}
