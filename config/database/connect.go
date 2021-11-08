package database

import (
	env "Jumia_todoList/config/environment"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	fmt.Printf("%#v\n", env.DatabaseEnv)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ",
		env.DatabaseEnv.Host, env.DatabaseEnv.UserName, env.DatabaseEnv.Password, env.DatabaseEnv.Name, env.DatabaseEnv.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(errors.Errorf("Failed to Intialize Db connection %s", err))
	}
	return db
}

func Migrate(orm *gorm.DB, entities ...interface{}) {
	orm.Migrator().AutoMigrate(entities...)
}
