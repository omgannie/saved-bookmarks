package main

import (
	"github.com/gin-gonic/gin"
)

type bookmark struct {
	ID          string `json:"id"`
	link_text   string
	href        string
	description string
}

func main() {
	router := gin.Default()
	models.connectDB()
	
	router.GET("/bookmarks", getBookmarks)

	router.Run("localhost:8080")
}

func getBookmarks(c *gin.Context) {
	// bookmarks := allBookmarks()
	// c.IndentedJSON(http.StatusOK, bookmarks)
}
