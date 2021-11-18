package env

import (
	"fmt"

	"github.com/spf13/viper"
)

var DatabaseEnv Database

type Database struct {
	UserName string
	Password string
	Name     string
	Port     string
	Host     string
}

func (d Database) must() error {

	if d.Host == "" || d.Port == "" || d.Name == "" || d.UserName == "" || d.Password == "" {
		panic(fmt.Errorf("MISSING DATABASE CONNECTION ENVIREOMENT"))
	}
	return nil
}

func (d *Database) load() {
	d.UserName = viper.Get("APP_DB_USERNAME").(string)
	d.Password = viper.Get("APP_DB_PASSWORD").(string)
	d.Name = viper.Get("APP_DB_NAME").(string)
	d.Port = viper.Get("APP_DB_PORT").(string)
	d.Host = viper.Get("APP_DB_HOST").(string)
}

func (d Database) export() {
	DatabaseEnv = d
}
