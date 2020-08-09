package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Config *DatabaseConfig
	db     *sql.DB
}

func (d *Database) Init(config *DatabaseConfig) error {
	if config == nil {
		return fmt.Errorf("nil DatabaseConfig")
	}
	d.Config = config
	return nil
}

func (d *Database) Connect() error {
	var err error
	d.db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", d.Config.User, d.Config.Password, d.Config.Host, d.Config.Database))
	if err != nil {
		return err
	}
	return nil
}
