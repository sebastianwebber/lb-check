package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

func checkIsReady(keyMetric string) error {

	cmd := exec.Command(viper.GetString("pg-isready-bin"), buildArgs()...)
	log.Printf("CMD: %#v ARGS: %#v\n", cmd.Path, cmd.Args)

	err := cmd.Run()
	if err != nil {
		return monitorError(keyMetric, fmt.Errorf("could not run cmd: %w", err))
	}

	return monitorError(keyMetric, nil)
}

func buildArgs() []string {
	var out []string

	if db.Options().Password != "" {
		os.Setenv("PGPASSWORD", db.Options().Password)
	}

	if viper.GetString("host") != "" {
		out = append(out, fmt.Sprintf("--host=%s", viper.GetString("host")))
	}
	out = append(out, fmt.Sprintf("--port=%d", viper.GetInt("port")))

	out = append(out, fmt.Sprintf("--dbname=%s", viper.GetString("dbname")))
	out = append(out, fmt.Sprintf("--username=%s", viper.GetString("username")))

	return out
}
