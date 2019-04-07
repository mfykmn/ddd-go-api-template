package db

import (
	"database/sql"

	"github.com/mafuyuk/ddd-go-api-template/config"

	"github.com/go-sql-driver/mysql"
)

type Client struct {
	*sql.DB
}

func NewMySQL(config *config.DB) (*Client, error) {
	conf := &mysql.Config{
		User:                 config.User,
		Passwd:               config.Password,
		Addr:                 config.Host + config.Port,
		Net:                  "tcp",
		DBName:               config.DBName,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &Client{db}, nil
}

func (c *Client) Close() {
	c.DB.Close()
}
