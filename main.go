package main

import (
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App){
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(PostLead)
	app.Delete(DeleteLead)
}

func main(){
	app := fiber.New()
	setupRoutes(app)
	app.listen(3000)  // Listens to Port 3000
}