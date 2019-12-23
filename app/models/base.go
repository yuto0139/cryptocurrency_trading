package models

import (
	"cryptocurrency_trading/config"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	tableNameSignalEvents = "signal_events"
)

// DbConnection ...
var DbConnection *sql.DB

// GetCandleTableName ...
func GetCandleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", productCode, duration)
}

func init() {
	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	createTable := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
	          time DATETIME PRIMARY KEY NOT NULL,
	          product_code VARCHAR(255),
	          side VARCHAR(255),
	          price FLOAT,
	          size FLOAT)`, tableNameSignalEvents)
	_, err = DbConnection.Exec(createTable)
	if err != nil {
		log.Fatalln(err)
	}

	for _, duration := range config.Config.Durations {
		tableName := GetCandleTableName(config.Config.ProductCode, duration)
		c := fmt.Sprintf(`
	          CREATE TABLE IF NOT EXISTS %s (
	          time DATETIME PRIMARY KEY NOT NULL,
	          open FLOAT,
	          close FLOAT,
	          high FLOAT,
	          low FLOAT,
						volume FLOAT)`, tableName)
		_, err = DbConnection.Exec(c)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
