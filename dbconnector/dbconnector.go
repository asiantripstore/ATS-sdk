package dbconnector

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var userDB *gorm.DB

//InitDB init connection to the postgres DB
func InitDB(dsn string) {
	var err error
	userDB, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}))
	if err != nil {
		log.WithField("InitDB", "OpenGorm").Error(err)
	}
}

//CheckDB check connection and GetDB
func CheckDB() (*gorm.DB, error) {
	db, err := userDB.DB()
	if err != nil {
		log.WithField("CheckDB", "Access DB").Error(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.WithField("CheckDB", "First Ping failed").Error(err)
		err = db.Ping()
		if err != nil {
			log.WithField("CheckDB", "Second Ping failed, connection to DB not possible").Error(err)
			return nil, err
		}
	}
	return userDB, nil
}
