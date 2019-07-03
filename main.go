package main

import (
	"./cmd"
	"os"
	"os/user"
)

type Config struct {
	user string
	last_played string
	current_host string
}

func checkIfSingHasRunBefore() bool {
	if _, err := os.Stat(".sing"); os.IsNotExist(err) {
		return false
	} else { return true }
}


func checkIfConfigPresent() bool{
	if _, err := os.Stat(".sing/config"); os.IsNotExist(err) {
		return false
	} else { return true }
}

func readConfigFromFile () Config{
	u, err := user.Current()
	checkErr(err)
	print(u.Name)
	return Config{}
}

func wtiteConfigToFile (configFile *os.File, config Config) {
	configFile.Write([]byte(`uname="` + config.user + `"` + "\n"))
	configFile.Write([]byte(`lastplayed="` + config.last_played + `"` + "\n"))
	configFile.Write([]byte(`currenthost="` + config.current_host + `"` + "\n"))
}

func setupConfig () {
	if err := os.Mkdir(".sing", os.ModeDir); err != nil {
		os.Chdir(".sing")
		if configFile, err := os.Create("config"); err == nil {
			if u, err := user.Current(); err == nil {
				blankConfig := Config{
					user: u.Name,
					last_played: "-",
					current_host: "-",
				}
				wtiteConfigToFile(configFile, blankConfig)
			}
		}
	}
}

func checkErr (e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	cmd.Execute()
}