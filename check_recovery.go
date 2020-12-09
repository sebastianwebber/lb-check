package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func checkRecovery() error {
	var out struct {
		Recovery bool
	}

	_, err := db.QueryOne(&out, viper.GetString("query-recovery"))

	if err != nil {
		return fmt.Errorf("could not compute delay bytes: %w", err)
	}

	log.Printf("Recovery %v\n", out)

	if !out.Recovery {
		return errors.New("Postgres Instance isn't in recovery")
	}

	return nil
}
