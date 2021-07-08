package lib

import (
	"fmt"

	"git.naspersclassifieds.com/olx/olxpeople/e2e/api-revenue-manager/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

//Database struct to get client and status
type Database struct {
	Connection *gorm.DB
	Status     bool
}

//DB variable to get Database instance
var DB Database

func init() {
	var err error
	if DB.Connection, err = GetConnection(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Connected to databases")
	}
}

// GetConnection : This can be called from anywhere in the code to get the DB connection
func GetConnection() (dbConnection *gorm.DB, err error) {
	if DB.Status == true {
		return DB.Connection, nil
	}

	dbSettings := settings.GetSettings().Database

	var conn *gorm.DB
	conn, err = gorm.Open(postgres.Open(dbSettings.ConnectionString()), &gorm.Config{
		Logger:               gormLogger.Default.LogMode(gormLogger.Error),
		DisableAutomaticPing: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "revenue.",
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("failed to connect to database", "error", err)
		return nil, err
	} else {
		_, err := conn.DB()
		if err != nil {
			panic(err)
		}
		DB.Status = true
		return conn, nil
	}
}
