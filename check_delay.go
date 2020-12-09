package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/spf13/viper"
)

func checkDelayBytes() error {
	var out struct {
		Delay uint64
	}

	_, err := db.QueryOne(&out, viper.GetString("query-delay-bytes"))

	if err != nil {
		return fmt.Errorf("could not compute delay bytes: %w", err)
	}

	log.Printf("Log bytes: %s", humanize.Bytes(out.Delay))

	return nil
}

func checkDelayDuration() error {
	var out struct {
		Delay time.Duration
	}

	_, err := db.QueryOne(&out, viper.GetString("query-delay-duration"))

	if err != nil {
		return fmt.Errorf("could not compute delay bytes: %w", err)
	}

	log.Printf("Log bytes: %v", out.Delay)

	return nil
}
