package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kabandr/go-fiber-api/controllers"
)

func CatchphrasesRouter(route fiber.Router) {
	route.Get("/", controllers.GetAllCatchphrases)
	route.Get("/:id", controllers.GetCatchphrases)
	route.Post("/", controllers.AddCatchphrase)
	route.Put("/:id", controllers.UpdateCatchphrase)
	route.Delete("/:id", controllers.DeleteCatchphrase)
	route.Delete("/", controllers.DeleteAllCatchphrases)
}
