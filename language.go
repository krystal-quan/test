package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Language int

const (
	Vietnamese Language = 0
	English    Language = 1
)

func (lan *Language) String() string {

	switch *lan {
	case Vietnamese:
		return "Vietnamese"
	case English:
		return "English"
	default:
		return "UnSpecified"
	}
}

// func (lan *Language) loadContent() string {

// }

// type Command struct {
// 	Name   string
// 	Detail map[string]string
// }
// type Command map[string]string

type Command map[string]string
type Content [2]Command

func (content *Content) setup() {
	rawData, _ := ioutil.ReadFile("Vn.json")
	err := json.Unmarshal(rawData, &content[0])
	if err != nil {
		log.Println(err)
	}

	rawData, _ = ioutil.ReadFile("Eng.json")
	err = json.Unmarshal(rawData, &content[1])
	if err != nil {
		log.Println(err)
	}

}

func (content *Content) loadContent(lan Language, command string, detail string) string {
	return content[lan][command+detail]
}
