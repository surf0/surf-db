package main

import (
	"server/controllers"

    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/records/map/:map", getRecordsByMap)
	router.Run("localhost:8083")
}

func getRecordsByMap(c *gin.Context) {
    mapName := c.Param("map")

	record := controllers.GetRecordsByMapName(mapName)

	if record == nil {    
 	   c.JSON(http.StatusNotFound, gin.H{"error": "map not found"})
	} else {
		c.JSON(http.StatusOK, record)
	}
}