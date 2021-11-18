package database

import (
	env "Jumia_todoList/config/environment"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(e env.Database) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ",
		e.Host, e.UserName, e.Password, e.Name, e.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(errors.Errorf("Failed to Initialize Db connection %s", err))
	}
	return db
}

func Migrate(orm *gorm.DB, entities ...interface{}) {
	orm.Migrator().AutoMigrate(entities...)
}
