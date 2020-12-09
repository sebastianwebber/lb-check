package main

import (
	"log"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/go-pg/pg"
	"github.com/gofiber/fiber/v2"
)

type checkItem struct {
	HelpMsg   string
	CheckFunc func(string) error
}

var checkList = map[string]checkItem{
	"pg_isready":          {HelpMsg: "Checks if postgres is running", CheckFunc: checkIsReady},
	"query-delay-seconds": {HelpMsg: "Checks the replica delay in seconds", CheckFunc: checkDelaySeconds},
	"is_recovering":       {HelpMsg: "Checks if the replica is in recovery", CheckFunc: checkRecovery},
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
			err := opts.CheckFunc(checkName)

			if err != nil {
				c.Status(500)
				log.Println(err.Error())
				return c.JSON(map[string]string{
					"Error": err.Error(),
					"Check": checkName,
				})
			}

			checks[checkName] = "OK"
		}

		c.Status(200)
		return c.JSON(map[string]interface{}{
			"OK":          "all checks ran without error",
			"Validations": checks,
		})
	})

	err := app.Listen(":3000")

	if err != nil {
		log.Printf("could not start the webserver: %s", err.Error())
	}
}
