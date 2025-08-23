package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type event struct {
    ID     string  `json:"id"`
    Title string `json:"title"`
	Description string `json:"description"`
	Date string `json:"date"`
	StartTime string `json:"start_time"`
	Duration string `json:"duration"`
	Location string `json:"location"`
	MaxCapacity int `json:"max_capacity"`
	Price float64 `json:"price"`
	Category string `json:"category"`
	Tags []string `json:"tags"`
}

var events = []event{
	{ID:"1", Title: "Event 1", Description: "Description 1", Date: "Date 1", StartTime: "StartTime 1", Duration: "Duration 1", Location: "Location 1", MaxCapacity: 1, Price: 1.0, Category: "Category 1", Tags: []string{"Tag 1"}},
	{ID:"2", Title: "Event 2", Description: "Description 2", Date: "Date 2", StartTime: "StartTime 2", Duration: "Duration 2", Location: "Location 2", MaxCapacity: 2, Price: 2.0, Category: "Category 2", Tags: []string{"Tag 2"}},

}


func getEvents(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, events)
}

func main() {
	fmt.Println("Hello World")
	router := gin.Default()

	// We may remove this one after we implement nginx or a proxy server.
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.GET("/events", getEvents)
	router.Run()
}