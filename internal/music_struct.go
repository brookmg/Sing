package internal

import (
	"encoding/json"
	"strings"
)

type MusicStruct struct {
	Fname string `json:"filename"`
	Fpath string `json:"filepath"`
}

type MusicMoreDetail struct {
	Artist string `json:"artist"`
	Title string `json:"title"`
	Album string `json:"album"`
	Year int `json:"year"`
	Genre string `json:"genre"`
}

type MusicFiles []MusicStruct

func (s MusicFiles) Len () int {
	return len(s)
}

func (s MusicFiles) Less (i , j int) bool {
	return s[i].Fname < s[j].Fname
}

func (s MusicFiles) Swap (i , j int) {
	s[i] , s[j] = s[j] , s[i]
}

func (s MusicFiles) Search (q string) MusicFiles{
	returnable := MusicFiles{}
	for i := range s {
		if strings.Contains(s[i].Fname , q) || strings.Contains(s[i].Fpath , q) {
			returnable = append(returnable , s[i])
		}
	}
	return returnable
}

func (s MusicFiles) ToJson () string {
	filesJson, _ := json.Marshal(s)
	return string(filesJson)
}

func (s MusicMoreDetail) ToJson () string {
	detailJson, _ := json.Marshal(s)
	return string(detailJson)
}