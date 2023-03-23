package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/wokacz/hermod/model"
	"github.com/wokacz/hermod/pkg/database"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/__health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	app.Use("/ws", func(ctx *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(ctx) {
			ctx.Locals("allowed", true)
			return ctx.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))
}

func init() {
	// Initialize database connection.
	err := database.Init()
	if err != nil {
		panic(err.Error())
	}
	// Migrate database.
	err = database.DB.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.AutoMigrate(&model.Notification{})
		return err
	})
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	err := app.Listen(":3015")
	if err != nil {
		panic(err.Error())
	}
}
