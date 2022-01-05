package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/kingdevnl/GoLive"
	"io/ioutil"
	"log"
)

type ChildComponent struct {
	GoLive.Component
}

func (c ChildComponent) OnMount(state *GoLive.State, args []interface{}) {
	state.Set("counter", 0)
	state.Set("message", "This is a paragraph.")
}

//func (c ChildComponent) OnEvent(event string, name string, data GoLive.Map, state *GoLive.State) {
//	log.Println("ChildOnEvent:", event, name)
//	if event == "click" {
//		counter := state.Get("counter").(int)
//		counter += 1
//		state.Set("counter", counter)
//		c.ReRender(state)
//	}
//}

func main() {

	GoLive.RegisterComponent("MyComponent", &MyComponent{})
	GoLive.RegisterComponent("ChildComponent", &ChildComponent{})

	engine := html.New("./main/views", ".tpl")

	engine.Reload(true)

	GoLive.SetupEngine(engine)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/golive.js", func(ctx *fiber.Ctx) error {
		buff, err := ioutil.ReadFile("./js/dist/golive.js")
		if err != nil {
			log.Println(err)
			return ctx.Status(500).SendString("Internal Server Error")
		}
		ctx.Set("Content-Type", "application/javascript")
		return ctx.SendString(string(buff))
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", nil)
	})

	GoLive.SetupHandler(GoLive.Config{Path: "/ws"}, app)

	//GoLive.SetupGarbageCollector()

	GoLive.InitComponents()

	log.Fatal(app.Listen(":3000"))
}
