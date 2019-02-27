package internal

import (
	"encoding/json"
	"fmt"
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
	q = strings.ToLower(q)
	qSpliced := splitBySpace(q)
	returnable := MusicFiles{}
	tempReturnable := MusicFiles{}

	// find all the items that contain the first element in the splited query
	for i := range s {
		if strings.Contains(strings.ToLower(s[i].Fpath) , qSpliced[0]) || strings.Contains(strings.ToLower(s[i].Fname) , qSpliced[0]) {
			returnable = append(returnable , s[i])
		}
	}

	for qi := range qSpliced {
		tempReturnable = returnable
		for i := range returnable {
			//if qi == 0 || i == 0 { continue }
			if !strings.Contains(strings.ToLower(returnable[i].Fpath) , qSpliced[qi]) && !strings.Contains(strings.ToLower(returnable[i].Fname) , qSpliced[qi]) {
				tempReturnable = removeItem(tempReturnable, returnable[i])
			}
		}
		returnable = tempReturnable
	}

	return returnable
}

func (s MusicFiles) SearchStrict (q string) MusicFiles{
	q = strings.ToLower(q)
	fmt.Print(q)
	returnable := MusicFiles{}
	for i := range s {
		if strings.Contains(strings.ToLower(s[i].Fname), q) || strings.Contains(strings.ToLower(s[i].Fpath), q) {
			returnable = append(returnable, s[i])
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

func removeItem (slice MusicFiles , item MusicStruct) MusicFiles {
	returnable := MusicFiles{}
	for i := range slice {
		if slice[i] != item {
			returnable = append(returnable, slice[i])
		}
	}
	return returnable
}