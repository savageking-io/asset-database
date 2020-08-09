package main

import (
	log "github.com/sirupsen/logrus"
)

func ServiceStart() error {
	log.Infof("Starting Asset-Database service")
	config := new(Config)
	if err := config.Init("config.yaml"); err != nil {
		log.Errorf("Failed to read configuration: %s", err.Error())
		return err
	}
	db := new(Database)
	if err := db.Init(config.Database); err != nil {
		log.Errorf("Failed to initialize database: %s", err.Error())
		return err
	}
	if err := db.Connect(); err != nil {
		log.Errorf("Failed to connect to database: %s", err.Error())
		return err
	}
	log.Infof("Initialization complete. Running...")
	return nil
}
