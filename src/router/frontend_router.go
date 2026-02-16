package router

import "github.com/gofiber/fiber/v2"

func FrontendRoutes(app *fiber.App) {
	app.Static("/frontend", "./frontend")
	app.Static("/uploads", "./frontend/uploads")

	app.Get("/", func(c *fiber.Ctx) error { 
		return c.SendFile("./frontend/index.html") 
	})

	app.Get("/admin.html", func(c *fiber.Ctx) error { 
		return c.SendFile("./frontend/admin.html") 
	})

	app.Get("/admin", func(c *fiber.Ctx) error { 
		return c.SendFile("./frontend/admin.html") 
	})

	app.Get("/login", func(c *fiber.Ctx) error { 
		return c.SendFile("./frontend/login.html") 
	})

	app.Get("/portfolio", func(c *fiber.Ctx) error { 
		return c.SendFile("./frontend/index.html") 
	})
}
