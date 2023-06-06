package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

type Verse string

type Chapter map[string]string

type Book map[string]map[string]string

type Bible map[string]map[string]map[string]string

type SearchReslut struct {
	Result []Verse
}

func main() {

	router := gin.Default()

	router.GET("/bible/:version", getBible)
	router.GET("/booklist/:version", getBookList)
	router.GET("/:book/:version", getBook)

	router.Run("localhost:8080")
}

// Return the whole Bible for a given version.
func getBible(c *gin.Context) {
	// Make sure the version is supported.
	version := titleCase(c.Param("version"))
	if isVersionSupported(version) {
		bible, err := loadBible(version)

		// If an error occured while loading the Bible, we stop.
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// Return the Bible content.
		c.IndentedJSON(http.StatusOK, bible)

	} else {
		//Return an error for unsupported versions
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Version not supported"})
		return
	}
}

// Returns the list of book names in a given Bible version.
func getBookList(c *gin.Context) {
	// Make sure the version is supported.
	version := titleCase(c.Param("version"))
	if isVersionSupported(version) {
		bible, err := loadBible(version)

		// If an error occured while loading the Bible, we stop.
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// Using the Mapkeys function to get a slice of the keys
		// which are the books

		booklist := make([]string, 0)

		for k, _ := range bible {
			booklist = append(booklist, k)
		}

		c.IndentedJSON(http.StatusOK, booklist)

	} else {
		//Return an error for unsupported versions
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Version not supported"})
		return
	}
}

// Returns the content of a Bible book.
func getBook(c *gin.Context) {
	book := titleCase(c.Param("book"))
	version := titleCase(c.Param("version"))

	// Make sure the version is supported.
	if isVersionSupported(version) {
		bible, err := loadBible(version)

		// If an error occured while loading the Bible, we stop.
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// Check if the provided book name exists in the specified version.
		_, bookExist := bible[book]
		if bookExist {
			// Return the requested Book.
			c.IndentedJSON(http.StatusOK, bible[book])

		} else {
			// Return an error if book not found.
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
			return
		}

	} else {
		//Return an error for unsupported versions.
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Version not supported"})
		return
	}

}

// Return strings in Title case.
func titleCase(s string) string {
	titleCase := cases.Title(language.Und)
	return titleCase.String(s)
}

// Loads the bible file and unmarshal the content into a map.
func loadBible(version string) (Bible, error) {
	// Path to the Bible file.
	data_folder := "../data/"
	data_extension := ".yml"

	// Path to the Bible file.
	bible_file_path := data_folder + version + data_extension

	// Read the Bible content.
	bible_content, err := os.ReadFile(bible_file_path)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("bible file not found")
	}

	//Unmarshal the Bible YAML data
	var bible Bible
	err = yaml.Unmarshal(bible_content, &bible)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("cannot unmarshal Bible content")
	}

	return bible, nil
}

// Checks if a version is supported.
func isVersionSupported(version string) bool {
	// The format is #<version1>#<version2>#.
	// We don't want to use a slice to hols the list of supported version.
	supported_versions := "#kjv#lsg#test#"

	// This condition checks if the version is among the supported ones.
	// "#"+version+"#" is a tip to eliminate cases where "version" is empty.
	return strings.Contains(supported_versions, "#"+version+"#")
}

// TODO:
//		- GET /:book/:chapter/:version
//		- GET /:book/:chapter/:verse/:version
//		- GET /versions
