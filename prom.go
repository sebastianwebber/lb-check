package main

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var metrics = map[string]prometheus.Gauge{}

func registerProm() {
	for checkName, opts := range checkList {
		metrics[checkName] = promauto.NewGauge(
			prometheus.GaugeOpts{
				Name: strings.ReplaceAll(
					fmt.Sprintf(
						"%s_%s_failing",
						appName,
						checkName,
					),
					"-",
					"_",
				),
				Help: opts.HelpMsg,
			})
	}
}

func monitorError(keyMetric string, err error) error {

	if err != nil {
		metrics[keyMetric].Set(1)
		return err
	}

	metrics[keyMetric].Set(0)
	return nil
}
