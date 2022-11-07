package main

import (
	"server/controllers"
	_ "server/models"

    "net/http"

    "github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "server/docs"
)

// @title Surf DB API
// @version 1.0
// @description CS:GO & CS:S surf record and map data. Github: https://github.com/surf0/surf-db
// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1") 
	{
		sh := v1.Group("/sh") 
		{
			sh.GET("/records/map/:map", getRecordsByMapSH)
		}
	}
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("localhost:8080")
}

// @Summary get all records for map by map name
// @Tags surfheaven
// @Produce json
// @Param mapname path string true "map name"
// @Success 200 {array} models.Record
// @Router /sh/records/map/{mapname} [get]
func getRecordsByMapSH(c *gin.Context) {
    mapName := c.Param("map")

	record := controllers.GetRecordsByMapNameSH(mapName)

	if record == nil {    
 	   c.JSON(http.StatusNotFound, gin.H{"error": "map not found"})
	} else {
		c.JSON(http.StatusOK, record)
	}
}