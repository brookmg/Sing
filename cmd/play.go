package cmd

import (
	"../internal"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

var random bool
var genre string
var contain string

func init() {
	playCmd.Flags().BoolVarP(&random, "random" , "r" , false , "Pick a random song from the indexed list")
	playCmd.Flags().StringVarP(&genre , "genre" , "g" , "" , "Genre for the chosen music")
	playCmd.Flags().StringVarP(&contain, "contain" , "c", "", "String that could be inside the name of the music")
	rootCmd.AddCommand(playCmd)
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Command to play music",
	Long:  `Command to play offline music on the device, might be random or with specific tag`,
	TraverseChildren:true,
	Run: func(cmd *cobra.Command, args []string) {
		musicList := internal.GetMusicList(internal.AvailableFileSystems())
		if random {
			if contain != "" {
				musicList = musicList.Search(contain)
			}

			if musicList.Len() > 0 {
				rand.Seed(time.Now().UnixNano())
				var r= rand.Intn(musicList.Len())
				var item= musicList[r].Fpath + "\\" + musicList[r].Fname
				println("Playing `" + musicList[r].Fname + "` Now.")
				item = strings.Replace(item, "\\", "/", -1)
				ex := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", item)
				_, err := ex.Output()

				if err != nil {
					log.Fatalf(err.Error())
				}
			} else {
				println("We couldn't find any music.")
			}
		}
	},
}
