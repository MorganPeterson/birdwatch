package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var frontend embed.FS

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

func postEntry(c *gin.Context) {
	var newEntry Entry

	if err := c.ShouldBindJSON(&newEntry); err != nil {
		log.Printf("birdwatch: Entry Post: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
	err := connectDatabase()
	if err != nil {
		log.Fatalf("Birdwatch: Database Connection: %s\n", err.Error())
	}

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	v1 := router.Group("/api/v1")
	v1.GET("/entry", getEntry)
	v1.POST("/entry", postEntry)

	router.Run("localhost:8080")
}
