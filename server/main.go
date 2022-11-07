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
// @tag.name surfheaven
// @tag.description surfheaven.eu CS:GO surf servers
func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1") 
	{
		sh := v1.Group("/sh") 
		{
			sh.GET("/records/map/:map", getRecordsByMapSH)
			sh.GET("/records/stage/:map/:track", getRecordsByStageSH)
			sh.GET("/records/stage/:map/", getRecordsStagesSH)
			sh.GET("/records/bonus/:map/:track", getRecordsByBonusSH)
			sh.GET("/records/bonus/:map/", getRecordsBonusesSH)
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

// @Summary get all records for specific stage
// @Tags surfheaven
// @Produce json
// @Param mapname path string true "map name"
// @Param stage path string true "stage e.g. 1, 2, ..."
// @Success 200 {array} models.Record
// @Router /sh/records/stage/{mapname}/{stage} [get]
func getRecordsByStageSH(c *gin.Context) {
    mapName := c.Param("map")
	track := c.Param("track")

	records := controllers.GetRecordsByStageSH(mapName, track)

	if records == nil {    
 	   c.JSON(http.StatusNotFound, gin.H{"error": "stage not found (only works for staged maps)"})
	} else {
		c.JSON(http.StatusOK, records)
	}
}

// @Summary get all records for all stages
// @Tags surfheaven
// @Produce json
// @Param mapname path string true "map name"
// @Success 200 {array} models.Record
// @Router /sh/records/stage/{mapname}/ [get]
func getRecordsStagesSH(c *gin.Context) {
    mapName := c.Param("map")

	records := controllers.GetRecordsStagesSH(mapName)

	if records == nil {    
 	   c.JSON(http.StatusNotFound, gin.H{"error": "stage not found (only works for staged maps)"})
	} else {
		c.JSON(http.StatusOK, records)
	}
}

// @Summary get all records for specific bonus
// @Tags surfheaven
// @Produce json
// @Param mapname path string true "map name"
// @Param bonus path string true "bonus e.g. 1, 2, ..."
// @Success 200 {array} models.Record
// @Router /sh/records/bonus/{mapname}/{bonus} [get]
func getRecordsByBonusSH(c *gin.Context) {
    mapName := c.Param("map")
	track := c.Param("track")

	records := controllers.GetRecordsByBonusSH(mapName, track)

	if records == nil {    
 	   c.JSON(http.StatusNotFound, gin.H{"error": "bonus not found"})
	} else {
		c.JSON(http.StatusOK, records)
	}
}

// @Summary get all records for all bonuses
// @Tags surfheaven
// @Produce json
// @Param mapname path string true "map name"
// @Success 200 {array} models.Record
// @Router /sh/records/bonus/{mapname}/ [get]
func getRecordsBonusesSH(c *gin.Context) {
    mapName := c.Param("map")

	records := controllers.GetRecordsBonusesSH(mapName)

	if records == nil {    
 	   c.JSON(http.StatusNotFound, gin.H{"error": "bonuses not found"})
	} else {
		c.JSON(http.StatusOK, records)
	}
}