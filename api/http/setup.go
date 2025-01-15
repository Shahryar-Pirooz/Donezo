package http

import (
	"donezo/app"
	"donezo/config"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run(cnf config.Config, app *app.App) {
	fiberApp := fiber.New()
	api := fiberApp.Group("/api/v1")
	registerGlobalRouts(api, app)
	registerProjectRouts(api, app)
	registerTaskRouts(api, app)

	log.Fatal(fiberApp.Listen(fmt.Sprintf("%s:%s", cnf.Server.Host, cnf.Server.Port)))
}

func registerTaskRouts(router fiber.Router, _ *app.App) {
	taskGroup := router.Group("/task")
	taskGroup.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "ok"})
	})

}

func registerProjectRouts(router fiber.Router, app *app.App) {
	panic("unimplemented")
}

func registerGlobalRouts(api fiber.Router, app *app.App) {
	panic("unimplemented")
}
