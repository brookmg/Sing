package internal

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func AvailableFileSystems() []string {
	return getDrives()
}

func getDrives() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		_, err := os.Open(string(drive)+":\\")
		if err == nil {
			r = append(r, string(drive))
		}
	}
	return r
}


func GetRefreshedMusicList(drives []string) MusicFiles {
	// start
	cmd := exec.Command("indexing.exe")
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	// wait or timeout
	donec := make(chan error, 1)
	go func() {
		donec <- cmd.Wait()
	}()
	select {
	case <-time.After(25 * time.Second):
		return MusicFiles{}
	case <-donec:
		os.Remove("allMusicFiles.csv")
		os.Rename("allMusicFiles.csv" , "allMusicFiles.csv")
		return GetMusicList(drives)
	}
}

func GetMusicList(drives []string) MusicFiles {
	csvFile, _ := os.Open("allMusicFiles.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var musicFiles MusicFiles
	for {
		line, e := reader.Read()
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatal(e)
		}

		if SliceContains(drives , string(strings.ToUpper(line[1])[0])) && strings.HasSuffix(line[0] , ".mp3") {
			musicFiles = append(musicFiles, MusicStruct{
				Fname: line[0],
				Fpath: line[1],
			})
		}

	}

	return musicFiles
}

func SliceContains (a []string , i string) bool {
	for _ , value := range a {
		if string(value) == i {
			return true
		}
	}
	return false
}

func removeCharactors (i string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	processed := reg.ReplaceAllString(i, "")
	return processed
}

func splitBySpace (i string) []string {
	return strings.Fields(i)
}

func FileExt (filename string) string {
	if len(filename) == 0 {
		return ""
	}

	dotIndex := strings.LastIndex(filename , ".")
	return filename[dotIndex:]
}
