package main

import (
	"log"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/go-pg/pg"
	"github.com/gofiber/fiber/v2"
)

type checkItem struct {
	HelpMsg   string
	CheckFunc func() error
}

var checkList = map[string]checkItem{
	"pg_isready": {HelpMsg: "postgres is running?", CheckFunc: checkIsReady},
	// "replica_delay_bytes":    {HelpMsg: "Replica delay in bytes", CheckFunc: checkDelayBytes},
	"replica_delay_duration": {HelpMsg: "Replica delay interval", CheckFunc: checkDelayDuration},
	"is_recovering":          {HelpMsg: "Replica is in recovery?", CheckFunc: checkRecovery},
}

func init() {
	registerProm()
}

func main() {

	db = pg.Connect(buildConn())
	defer db.Close()

	app := fiber.New()

	prometheus := fiberprometheus.New(appName)
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	app.Get("/check", func(c *fiber.Ctx) error {

		checks := map[string]string{}

		for checkName, opts := range checkList {
			err := opts.CheckFunc()

			if err != nil {
				c.Status(500)
				log.Println(err.Error())
				return c.JSON(map[string]string{
					"Error": err.Error(),
					"Check": checkName,
				})
			}

			checks[checkName] = opts.HelpMsg
		}

		c.Status(200)
		return c.JSON(map[string]interface{}{
			"OK":          "all good.",
			"Validations": checks,
		})
	})

	err := app.Listen(":3000")

	if err != nil {
		log.Printf("could not start the webserver: %s", err.Error())
	}
}
