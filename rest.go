package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Rest struct {
	Config *RestConfig
	User   *User
}

func (d *Rest) Init(config *RestConfig, u *User) error {
	if config == nil {
		return fmt.Errorf("nil RestConfig")
	}
	if u == nil {
		return fmt.Errorf("nil User")
	}
	d.Config = config
	d.User = u

	http.HandleFunc("/user", d.HandleUser)
	return nil
}

func (d *Rest) HandleUser(w http.ResponseWriter, r *http.Request) {
	if d.User == nil {
		log.Errorf("Can't handle /user request due to nil User")
		return
	}
}
