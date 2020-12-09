package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

func checkDelaySeconds() error {
	var out struct {
		Delay time.Duration
	}

	_, err := db.QueryOne(&out, viper.GetString("query-delay-seconds"))

	if err != nil {
		return fmt.Errorf("could not compute delay bytes: %w", err)
	}

	log.Printf("Log bytes: %v", out.Delay)

	return nil
}
