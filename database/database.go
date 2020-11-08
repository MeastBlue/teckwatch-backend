package database

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/meastblue/teckwatch/config"
)

func New() (*sqlx.DB, error) {
	conf, err := config.InitConfig()
	if err != nil {
		return nil, err
	}

	dbc := conf.GetDatabaseConf()
	if dbc.Host == "" || dbc.Port == "" || dbc.Name == "" || dbc.User == "" || dbc.Pass == "" {
		err := errors.New("One of database info is empty")
		return nil, err
	}

	db, err := connectDatabase(dbc)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectDatabase(c *config.Database) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User,
		c.Pass, c.Name)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
