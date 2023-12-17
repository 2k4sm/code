package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(
			fiber.Map{
				"Congrats!": "Yay!...You've Created a JSON API using Fiber...",
			},
		)
	})

	app.Get("/randomfacts", randomfacts)

	log.Fatal(app.Listen(":8080"))
}

func randomfacts(ctx *fiber.Ctx) error {
	facts := []string{
		"The circulatory system is more than 60,000 miles long",
		"Sudan has more pyramids than any country in the world",
		"The bumblebee bat is the world’s smallest mammal",
		"The world’s first animated feature film was made in Argentina",
		"A one-way trip on the Trans-Siberian Railway involves crossing 3,901 bridges",
		"The Golden Girls was supposed to have a different theme song",
		"There’s enough gold inside Earth to coat the planet",
		"Cleveland was once the country’s fifth-largest city",
		"Japan has one vending machine for every 40 people",
	}

	fact := facts[rand.Intn(9)]

	return ctx.JSON(
		fiber.Map{
			"time": time.Now(),
			"fact": fact,
		},
	)
}
