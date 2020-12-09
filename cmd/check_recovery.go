package main

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

func checkRecovery(keyMetric string) error {

	var out struct {
		Recovery bool
	}

	_, err := db.QueryOne(&out, viper.GetString("query-recovery"))

	if err != nil {
		return monitorError(keyMetric, err)
	}

	log.Printf("Recovery %v\n", out)

	if !out.Recovery {
		return monitorError(keyMetric, errors.New("Postgres Instance isn't in recovery"))
	}

	return monitorError(keyMetric, nil)
}
