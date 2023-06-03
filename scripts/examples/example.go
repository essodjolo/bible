package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Bible is a nested struct that helps to load the Bible content.
type Bible struct {
	Books struct {
		// Will capture only the books I want to work with.
		Psalms map[string]map[string]string `yaml:"Psalms"`
		John   map[string]map[string]string `yaml:"John"`
		Jude   map[string]map[string]string `yaml:"Jude"`
	} `yaml:"books"`
}

func main() {
	// Path to the Bible file.
	const data = "../../data/kjv.yml"

	// Read the Bible content.
	kjv_bible, err := os.ReadFile(data)
	if err != nil {
		log.Fatal(err)
	}

	//Unmarshal the Bible YAML data into a Bible struct varibale
	var bible Bible
	err = yaml.Unmarshal(kjv_bible, &bible)
	if err != nil {
		log.Fatal(err)
	}

	// Printing Jean 3:16
	chapter := "3"
	verse := "16"
	fmt.Println("John " + chapter + ":" + verse)
	fmt.Println(bible.Books.John[chapter][verse])

	// Printing Psalms 23
	fmt.Println("\n\nPsalms 23")
	fmt.Println(bible.Books.Psalms["23"])

	// Printing Jude
	fmt.Println("\n\nJude:")
	fmt.Println(bible.Books.Jude)
}
