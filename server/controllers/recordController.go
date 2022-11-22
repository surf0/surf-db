package controllers

import (
	"server/ent"

	"context"
	"server/ent/recordsh"
	"server/ent/recordksf"
	"server/config"

	_ "github.com/go-sql-driver/mysql"

	"log"
)


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

func QueryRecordsStagesSH(mapName string) []*ent.RecordSh {
	var client *ent.Client = config.GetClient()
	records, err := client.RecordSh.
        Query().
        Where(
			recordsh.MapName(mapName),
			recordsh.Type("stage"),
		).
		Order(ent.Desc(recordsh.FieldTimestamp)).
        All(context.Background())
    if err != nil {
		log.Fatalf("sql error: %v", err)
        return nil
    }
    return records
}



func QueryRecordsByBonusSH(mapName string, track int) []*ent.RecordSh {
	var client *ent.Client = config.GetClient()
	records, err := client.RecordSh.
        Query().
        Where(
			recordsh.MapName(mapName),
			recordsh.Type("bonus"),
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

func QueryRecordsBonusesSH(mapName string) []*ent.RecordSh {
	var client *ent.Client = config.GetClient()
	records, err := client.RecordSh.
        Query().
        Where(
			recordsh.MapName(mapName),
			recordsh.Type("bonus"),
		).
		Order(ent.Desc(recordsh.FieldTimestamp)).
        All(context.Background())
    if err != nil {
		log.Fatalf("sql error: %v", err)
        return nil
    }
    return records
}

func QueryRecordsByMapNameKSF(mapName string) []*ent.RecordKsf {
	var client *ent.Client = config.GetClient()
	records, err := client.RecordKsf.
        Query().
        Where(
			recordksf.MapName(mapName),
		).
		Order(ent.Desc(recordksf.FieldTimestamp)).
        All(context.Background())
    if err != nil {
		log.Fatalf("sql error: %v", err)
        return nil
    }
    return records
}