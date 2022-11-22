package api

import (
	"server/controllers"

    "net/http"
	"strconv"
    "github.com/gin-gonic/gin"

	_ "server/docs"
)

// @Summary get all records for map by map name
// @Tags surfheaven
// @Produce json
// @Param mapname path string true "map name e.g. surf_summit"
// @Success 200 {array} ent.RecordSh
// @Router /sh/records/map/{mapname} [get]
func GetRecordsByMapSH(c *gin.Context) {
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
// @Param mapname path string true "map name e.g. surf_summit"
// @Param stage path string true "stage e.g. 1, 2, ..."
// @Success 200 {array} ent.RecordSh
// @Router /sh/records/stage/{mapname}/{stage} [get]
func GetRecordsByStageSH(c *gin.Context) {
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
// @Param mapname path string true "map name e.g. surf_summit"
// @Success 200 {array} ent.RecordSh
// @Router /sh/records/stage/{mapname}/ [get]
func GetRecordsStagesSH(c *gin.Context) {
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
// @Param mapname path string true "map name e.g. surf_summit"
// @Param bonus path string true "bonus e.g. 1, 2, ..."
// @Success 200 {array} ent.RecordSh
// @Router /sh/records/bonus/{mapname}/{bonus} [get]
func GetRecordsByBonusSH(c *gin.Context) {
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
// @Param mapname path string true "map name e.g. surf_summit"
// @Success 200 {array} ent.RecordSh
// @Router /sh/records/bonus/{mapname}/ [get]
func GetRecordsBonusesSH(c *gin.Context) {
    mapName := c.Param("map")

	records := controllers.QueryRecordsBonusesSH(mapName)

	if records == nil {    
 	   c.JSON(http.StatusNotFound, gin.H{"error": "bonuses not found"})
	} else {
		c.JSON(http.StatusOK, records)
	}
}

// @Summary get all records for map by map name
// @Tags ksf
// @Produce json
// @Param mapname path string true "map name e.g. surf_summit"
// @Success 200 {array} ent.RecordKsf
// @Router /ksf/records/map/{mapname} [get]
func GetRecordsByMapKSF(c *gin.Context) {
    mapName := c.Param("map")

	records := controllers.QueryRecordsByMapNameKSF(mapName)

	if records == nil {    
 	   c.JSON(http.StatusNotFound, gin.H{"error": "map not found"})
	} else {
		c.JSON(http.StatusOK, records)
	}
}