package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/keiwi/utils"

	// mysql db driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB abstraction
type DB struct {
	*gorm.DB
}

// NewMysqlDB - mysql database
func NewMysqlDB(user, password, host, port, dbname string) *DB {
	utils.Log.Debug("Initializing database connection")
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname))
	if err != nil {
		panic(err)
	}

	utils.Log.WithField("test", "foobar").Info("Database successfully connected")

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(false)

	utils.Log.Debug("Database connection initialization done")
	return &DB{db}
}
