package main

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
	"github.com/spf13/viper"
)

var db *pg.DB

func buildConn() *pg.Options {

	connOpts := &pg.Options{
		User:            viper.GetString("username"),
		Database:        viper.GetString("dbname"),
		PoolSize:        1,
		ApplicationName: appName,
	}

	if dbPassword != "" {
		connOpts.Password = dbPassword
	}

	if connOpts.Password == "" && os.Getenv("PGPASSWORD") != "" {
		connOpts.Password = os.Getenv("PGPASSWORD")
	}

	if viper.GetString("host") != "local socket" {
		connOpts.Network = "tcp"
		connOpts.Addr = fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("port"))
		return connOpts
	}

	connOpts.Network = "unix"

	return connOpts
}
