package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	HOST = "localhost"
	PORT = "8080"
)

// contains our embedded JavaScript code from vue
var frontend embed.FS

// getEntry is the high-level function to fetch data from sqlite and then
// return all rows to the frontend.
func getEntry(c *gin.Context) {
	entries, err := runSelect()
	if err != nil {
		log.Printf("birdwatch: DB query error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to query database"})
		return
	}

	if entries == nil {
		log.Print("birdwatch: No Records Found\n")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No Records Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": entries})
}

// postEntry is the high-level function for binding incoming object to Entry
// and then sending it off to sqlite. It returns `nil` on success and an error
// message otherwise.
func postEntry(c *gin.Context) {
	var newEntry Entry

	if err := c.ShouldBindJSON(&newEntry); err != nil {
		log.Printf("birdwatch: Entry Post: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// insert entry into the database
	success, err := runInsert(newEntry)
	if err != nil {
		log.Printf("birdwatch: DB insert error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error writing to database"})
	} else if success {
		c.JSON(http.StatusCreated, gin.H{"error": nil})
	} else {
		log.Printf("birdwatch: DB Unknown error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func main() {
	// connect to database. Full shutdown if not successful
	err := connectDatabase()
	if err != nil {
		log.Fatalf("Birdwatch: Database Connection: %s\n", err.Error())
	}

	address := fmt.Sprintf("%s:%s", HOST, PORT)

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	v1 := router.Group("/api/v1")
	// fetch all entries in the database
	v1.GET("/entry", getEntry)
	// post single entry
	v1.POST("/entry", postEntry)

	router.Run(address)
}
