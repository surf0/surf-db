package controllers

import (
	"server/models"
	"server/ent"

	"context"
	"server/ent/recordsh"
	"server/config"

	
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"log"

)

const dbUser = "db"
const dbPass = "password"
const dbName = "surfdb"

func QueryRecordsByMapNameSH(mapName string) []*ent.RecordSh {
	var client *ent.Client = config.GetClient()
	records, err := client.RecordSh.
        Query().
        Where(
			recordsh.MapName(mapName),
			recordsh.Type("map"),
		).
		Order(ent.Desc(recordsh.FieldTimestamp)).
        All(context.Background())
    if err != nil {
		log.Fatalf("sql error: %v", err)
        return nil
    }
    return records
}

func QueryRecordsByStageSH(mapName string, track int) []*ent.RecordSh {
	var client *ent.Client = config.GetClient()
	records, err := client.RecordSh.
        Query().
        Where(
			recordsh.MapName(mapName),
			recordsh.Type("stage"),
			recordsh.Track(track),
		).
		Order(ent.Desc(recordsh.FieldTimestamp)).
        All(context.Background())
    if err != nil {
		log.Fatalf("sql error: %v", err)
        return nil
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

func GetRecordsByBonusSH(mapName string, track string) []models.Record {
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)

	records := []models.Record{} 

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM records_sh WHERE map_name=? AND type='bonus' AND track=? ORDER BY timestamp DESC", mapName, track)
	
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

func GetRecordsBonusesSH(mapName string) []models.Record {
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)

	records := []models.Record{} 

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM records_sh WHERE map_name=? AND type='bonus' ORDER BY timestamp DESC", mapName)
	
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