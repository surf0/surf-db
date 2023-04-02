package main

import (
	"server/config"
	"server/api"


	"log"
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
// @tag.name ksf
// @tag.description KSF CS:S surf servers
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
			sh.GET("/records/map/:map", api.GetRecordsByMapSH)
			sh.GET("/records/stage/:map/:track", api.GetRecordsByStageSH)
			sh.GET("/records/stage/:map/", api.GetRecordsStagesSH)
			sh.GET("/records/bonus/:map/:track", api.GetRecordsByBonusSH)
			sh.GET("/records/bonus/:map/", api.GetRecordsBonusesSH)
		}

		ksf := v1.Group("/ksf")
		{
			ksf.GET("/records/map/:map", api.GetRecordsByMapKSF)
		}
	}
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("0.0.0.0:8080")
}
