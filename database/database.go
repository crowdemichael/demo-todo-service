package database

import (
	"database/sql"
	"fmt"

	"github.com/crowdeco/demo-todo-service/configs"
	"github.com/crowdeco/demo-todo-service/database/schema"
	driver "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"log"
)

var (
	Client   *sql.DB
	Database *gorm.DB
)

func Connect() {
	user := configs.Env.DbUser
	pass := configs.Env.DbPass
	host := configs.Env.DbHost
	port := configs.Env.DbPort
	name := configs.Env.DbName

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", user, pass, host, port, name)
	db, err := gorm.Open(driver.Open(conn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	Database = db
	Client, _ = db.DB()

	if configs.Env.Debug {
		schema.AutoMigrate(db)
	}
}
