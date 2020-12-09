package main

import (
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

func checkIsReady() error {
	cmd := exec.Command(viper.GetString("pg-isready-bin"))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not run cmd: %w", err)
	}
	return nil
}
