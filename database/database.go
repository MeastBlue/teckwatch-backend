package database

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

type DB struct {
	*sqlx.DB
}

func New(host, port, name, user, pass string) (*DB, error) {
	if host == "" || port == "" || name == "" || user == "" || pass == "" {
		err := errors.New("One of database info is empty")
		return nil, err
	}

	config := Config{
		Host: host,
		Port: port,
		Name: name,
		User: user,
		Pass: pass,
	}

	db, err := config.init()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (c *Config) init() (*DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User,
		c.Pass, c.Name)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	
	return &DB{db}, nil
}
