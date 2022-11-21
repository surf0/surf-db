package main

import (
	"server/controllers"
	"server/config"

    "net/http"
	"log"
	"strconv"
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

	client, err := config.NewClient()
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()

	if err != nil {
		log.Println("Failed to initialize client")
	}

	config.SetClient(client)

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

	records := controllers.QueryRecordsByMapNameSH(mapName)

	if records == nil {    
 	   c.JSON(http.StatusNotFound, gin.H{"error": "map not found"})
	} else {
		c.JSON(http.StatusOK, records)
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
	track, err := strconv.Atoi(c.Param("track"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "track/stage must be an integer"})
		return
	}

	records := controllers.QueryRecordsByStageSH(mapName, track)

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

	records := controllers.QueryRecordsStagesSH(mapName)

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
	track, err := strconv.Atoi(c.Param("track"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "track/bonus must be an integer"})
		return
	}

	records := controllers.QueryRecordsByBonusSH(mapName, track)

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

	records := controllers.QueryRecordsBonusesSH(mapName)

	if records == nil {    
 	   c.JSON(http.StatusNotFound, gin.H{"error": "bonuses not found"})
	} else {
		c.JSON(http.StatusOK, records)
	}
}