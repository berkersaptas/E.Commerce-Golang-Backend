package routes

import (
	"FluxStore/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/setUser", controllers.CreateUser)
	app.Post("/login", controllers.LoginUser)
	app.Post("/forgetPassword", controllers.ForgetPassword)
	app.Post("/resetPassword", controllers.ResetPassword)
}
