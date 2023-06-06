package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Bible map[string]map[string]map[string]string

func main() {
	// Path to the Bible file.
	data_folder := "data/"

	// Path to the Bible file.
	bible_file_path := data_folder + "kjv.yml"

	// Read the Bible content.
	bible_content, err := os.ReadFile(bible_file_path)
	if err != nil {
		log.Fatal(err)
	}

	//Unmarshal the Bible YAML data
	var bible Bible
	err = yaml.Unmarshal(bible_content, &bible)
	if err != nil {
		log.Fatal(err)
	}

	// Printing Jean 3:16
	chapter := "3"
	verse := "16"
	fmt.Println("John " + chapter + ":" + verse)
	fmt.Println(bible["John"][chapter][verse])

	// Printing Psalms 23
	fmt.Println("\n\nPsalms 23")
	fmt.Println(bible["Psalms"]["23"])

	// Printing Jude
	fmt.Println("\n\nJude:")
	fmt.Println(bible["Jude"])
}
