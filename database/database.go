package database

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"

	logging "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbConn *gorm.DB

const slowSqlThreshold = 200 * time.Millisecond

func CreateDBConn(uri string, logFullQuery bool) *gorm.DB {
	fmt.Println("Connecting to ", uri)

	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	dbConn, err = gorm.Open(postgres.Open(strings.TrimSpace(uri)), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logging.WithFields(logging.Fields{"error": err}).Error("db connection error")
		panic(err)
	}
	db, err := dbConn.DB()
	if err != nil {
		logging.WithFields(logging.Fields{"error": err}).Error("failed to access db")
		panic(err.Error())
	}
	// Apply the default settings
	db.SetMaxIdleConns(3 - 1/3)
	db.SetMaxOpenConns(3)
	db.SetConnMaxLifetime(time.Second * time.Duration(300))
	db.SetConnMaxIdleTime(time.Second * time.Duration(300))

	logging.Info("successfully created database connection")
	return dbConn
}

func GetDBConn() *gorm.DB {
	return dbConn
}
