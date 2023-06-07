package main

import (
	"errors"
	"fmt"
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

	// /versions is an alias for /booklist
	router.GET("/versionlist", getVersionList)
	router.GET("/versions", getVersionList)

	router.GET("/bible/:version", getBible)

	// /books is an alias for /booklist
	router.GET("/booklist/:version", getBookList)
	router.GET("/books/:version", getBookList)

	router.GET("/book/:bookname/:version", getBook)

	router.GET("/chapter/:bookname/:chapter/:version", getChapter)

	router.GET("/verse/:bookname/:chapter/:verse/:version", getVerse)

	router.Run("localhost:8080")
}

// Returns the whole Bible for a given version.
// Serves the GET /versionlist and GET /versions route.
func getVersionList(c *gin.Context) {
	_, versions_list := getSupportedVersions()
	c.IndentedJSON(http.StatusOK, versions_list)
}

// Returns the whole Bible for a given version.
// Serves the GET /bible/:version route.
func getBible(c *gin.Context) {
	// Make sure the version is supported.
	version := strings.ToLower(c.Param("version"))
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
// Serves the GET /books/:version and GET /booklist/:version routes.
func getBookList(c *gin.Context) {
	// Make sure the version is supported.
	version := strings.ToLower(c.Param("version"))
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
		//Returns an error for unsupported versions
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Version not supported"})
		return
	}
}

// Returns the content of a Bible book.
// Serves the GET /book/:bookname/:version route.
func getBook(c *gin.Context) {
	book := titleCase(c.Param("bookname"))
	version := strings.ToLower(c.Param("version"))

	// Make sure the version is supported.
	if isVersionSupported(version) {
		bible, err := loadBible(version)

		// If an error occured while loading the Bible, we stop.
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// Check if the provided book name exists in the specified version.
		_, bookExists := bible[book]
		if bookExists {
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

// Returns the content of a Bible Chapter.
// Serves the GET /chapter/:bookname/:chapter/:version route.
func getChapter(c *gin.Context) {
	book := titleCase(c.Param("bookname"))
	chapter := c.Param("chapter")
	version := strings.ToLower(c.Param("version"))

	// Make sure the version is supported.
	if isVersionSupported(version) {
		bible, err := loadBible(version)

		// If an error occured while loading the Bible, we stop.
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// Check if the provided book name exists in the specified version.
		_, bookExists := bible[book]
		if bookExists {

			_, chapterExists := bible[book][chapter]
			if chapterExists {
				// Return the requested Chapter.
				c.IndentedJSON(http.StatusOK, bible[book][chapter])

			} else {
				// Return an error if chapter not found.
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Chapter not found"})
				return
			}

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

// Returns the content of a Bible Verse.
// Serves the GET /verse/:bookname/:chapter/:verse/:version route.
func getVerse(c *gin.Context) {
	book := titleCase(c.Param("bookname"))
	chapter := c.Param("chapter")
	verse := c.Param("chapter")
	version := strings.ToLower(c.Param("version"))

	// Make sure the version is supported.
	if isVersionSupported(version) {
		bible, err := loadBible(version)

		// If an error occured while loading the Bible, we stop.
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// Check if the provided book name exists in the specified version.
		_, bookExists := bible[book]
		if bookExists {

			_, chapterExists := bible[book][chapter]
			if chapterExists {

				_, verseExists := bible[book][chapter][verse]
				if verseExists {

					// Return the requested Chapter.
					c.IndentedJSON(http.StatusOK, bible[book][chapter][verse])

				} else {
					// Return an error if verse not found.
					c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Verse not found"})
					return
				}

			} else {
				// Return an error if chapter not found.
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Chapter not found"})
				return
			}

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
	supported_versions, _ := getSupportedVersions()
	// This condition checks if the version is among the supported ones.
	// "#"+version+"#" is a tip to eliminate cases where "version" is empty.
	fmt.Println("#" + version + "#")
	return strings.Contains(supported_versions, "#"+version+"#")
}

// Return the supported version as string AND as a slice.
func getSupportedVersions() (string, []string) {
	// The format is #<version1>#<version2>#.
	// We don't want to use a slice to hols the list of supported version.
	versions_string := "#kjv#lsg#test#"
	versions_slice := strings.Split(versions_string, "#")
	// the first and last element are empty, we remove them.
	versions_slice = versions_slice[1 : len(versions_slice)-1]
	return versions_string, versions_slice
}

// Return strings in Title case.
func titleCase(s string) string {
	titleCase := cases.Title(language.Und)
	return titleCase.String(s)
}

// TODO:
//		- GET /:bookname/:chapter/:version
//		- GET /:bookname/:chapter/:verse/:version
//		- GET /search/:version/<keyword>
