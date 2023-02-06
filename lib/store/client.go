package store

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const (
	DefaultAppName        = "Default"
	DefaultConnectTimeout = 10 * time.Second
	DefaultPingTimeout    = 2 * time.Second
)

type DbConfig struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DbName string
	Addr   string
}

func NewMysqlConnection() (*sql.DB, error) {
	return (&DbConfig{
		User:   os.Getenv("DB_USER"),
		Pass:   os.Getenv("DB_PASS"),
		DbName: os.Getenv("DB_NAME"),
		Addr:   os.Getenv("DB_ADDR"),
	}).MysqlConnect()
}

func (c *DbConfig) MysqlConnect() (myc *sql.DB, err error) {
	cfg := mysql.Config{
		User:                 c.User,
		Passwd:               c.Pass,
		Net:                  "tcp",
		Addr:                 c.Addr,
		DBName:               c.DbName,
		AllowNativePasswords: true,
		Timeout:              DefaultConnectTimeout,
		ClientFoundRows:      true,
		MultiStatements:      true,
	}

	myc, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		err = errors.Wrap(err, "failed to create mysql client")
		return
	}

	myc.SetConnMaxLifetime(0)
	myc.SetMaxIdleConns(50)
	myc.SetMaxOpenConns(50)

	pingCtx, cancelPingCtx := context.WithTimeout(context.Background(), DefaultPingTimeout)
	defer cancelPingCtx()

	if err = myc.PingContext(pingCtx); err != nil {
		err = errors.Wrap(err, "failed to establish connection to mysql server")
		myc.Close()
	}

	return
}
