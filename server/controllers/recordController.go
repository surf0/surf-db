package controllers

import (
	"server/models"
	
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

const dbUser = "db"
const dbPass = "password"
const dbName = "surfdb"

func GetRecordsByMapNameSH(mapName string) []models.Record {
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)

	records := []models.Record{} 

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM records_sh WHERE map_name=? AND type='map' ORDER BY timestamp DESC", mapName)
	
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	
	for results.Next() {
		var record models.Record

		err = results.Scan(&record.Id, &record.Timestamp, &record.PlayerName, &record.PlayerId, &record.Type, &record.Track, &record.MapName, &record.Time, &record.Improvement, &record.Server)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, record)
	}

	return records
}

func GetRecordsByStageSH(mapName string, track string) []models.Record {
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)

	records := []models.Record{} 

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM records_sh WHERE map_name=? AND type='stage' AND track=? ORDER BY timestamp DESC", mapName, track)
	
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	
	for results.Next() {
		var record models.Record

		err = results.Scan(&record.Id, &record.Timestamp, &record.PlayerName, &record.PlayerId, &record.Type, &record.Track, &record.MapName, &record.Time, &record.Improvement, &record.Server)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, record)
	}

	return records
}

func GetRecordsStagesSH(mapName string) []models.Record {
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)

	records := []models.Record{} 

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM records_sh WHERE map_name=? AND type='stage' ORDER BY timestamp DESC", mapName)
	
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	
	for results.Next() {
		var record models.Record

		err = results.Scan(&record.Id, &record.Timestamp, &record.PlayerName, &record.PlayerId, &record.Type, &record.Track, &record.MapName, &record.Time, &record.Improvement, &record.Server)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, record)
	}

	return records
}