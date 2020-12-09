package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

func checkDelaySeconds(keyMetric string) error {

	var out struct {
		Delay int
	}

	_, err := db.QueryOne(&out, viper.GetString("query-delay-seconds"))

	if err != nil {
		return monitorError(keyMetric, fmt.Errorf("could not compute delay bytes: %w", err))
	}

	delay := time.Duration(out.Delay) * time.Second

	log.Printf("Replica delay: %v", delay)

	if delay > viper.GetDuration("max-delay") {
		return monitorError(keyMetric, fmt.Errorf("replica max delay achieved: %v behind the primary database", delay))
	}

	return monitorError(keyMetric, nil)
}
