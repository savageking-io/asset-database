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
	log.Infof("Initializing Database")
	db := new(Database)
	if err := db.Init(config.Database); err != nil {
		log.Errorf("Failed to initialize database: %s", err.Error())
		return err
	}
	log.Infof("Establishing a database connection")
	if err := db.Connect(); err != nil {
		log.Errorf("Failed to connect to database: %s", err.Error())
		return err
	}

	log.Infof("Initializing user subsystem")
	user := new(User)

	log.Infof("Initializing REST API subsystem")
	rest := new(Rest)
	if err := rest.Init(config.Rest, user); err != nil {
		log.Errorf("Failed to initialize REST: %s", err.Error())
		return err
	}

	log.Infof("Initialization complete. Running...")
	return nil
}
