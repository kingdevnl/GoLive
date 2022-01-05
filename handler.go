package GoLive

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

type Config struct {
	Path string `json:"path"`
}

func SetupHandler(config Config, app *fiber.App) {
	// Handler to handle all types if it's a websocket upgrade forward the request to the websocket handler
	app.Use(config.Path, func(ctx *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(ctx) {
			ctx.Locals("ctx", ctx)
			return ctx.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// Handler to handle websocket requests
	app.Get(config.Path, websocket.New(func(conn *websocket.Conn) {
		ws := NewWebsocket(conn)
		Connections[ws.ID] = ws
		log.Printf("WebSocket %s connected.\n", ws.ID)
		ws.Conn.SetCloseHandler(func(code int, text string) error {
			log.Printf("WebSocket %s disconnected.\n", ws.ID)
			ws.IsConnected = false

			// loop over all states
			for _id, state := range states {
				if state.SocketID == ws.ID {
					log.Println("Removing state:", _id)
					delete(states, _id)
				}
			}

			delete(Connections, ws.ID)
			return nil
		})

		for {
			var json Map

			if err := conn.ReadJSON(&json); err != nil {
				if !ws.IsConnected {
					return
				}
				log.Printf("Websocket %s read error: %s", ws.ID, err)
				break
			}

			_handleEvent(ws, json)

		}

	}))
}
