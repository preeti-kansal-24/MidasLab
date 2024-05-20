package database

import "preeti-kansal-24/MidasLab.git/schema"

func Migrate() {
	dbConn := GetDBConn()
	dbConn.AutoMigrate(&schema.UserProfile{}, &schema.Otps{})

}
