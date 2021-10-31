package env

import (
	"fmt"
	"github.com/spf13/viper"
)

var DatabaseEnv database

type database struct {
	UserName string
	Password string
	Name     string
	Port     string
	Host     string
}

func (d database) must() error {

	if d.Host == "" || d.Port == "" || d.Name == "" || d.UserName == "" || d.Password == "" {
		panic(fmt.Errorf("MISSING DATABASE CONNECTION ENVIREOMENT"))
	}
	return nil
}

func (d *database) load() {
	d.UserName = viper.Get("APP_DB_USERNAME").(string)
	d.Password = viper.Get("APP_DB_PASSWORD").(string)
	d.Name = viper.Get("APP_DB_NAME").(string)
	d.Port = viper.Get("APP_DB_PORT").(string)
	d.Host = viper.Get("APP_DB_HOST").(string)
}

func (d database) export() {
	DatabaseEnv = d
}
