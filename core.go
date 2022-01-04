package GoLive

import (
	"github.com/gofiber/template/html"
)

type Map map[string]interface{}
type EventFunc func(state *State, data Map)

var Engine *html.Engine
var Connections = make(map[string]*WebSocket)

func SetupEngine(engine *html.Engine) {
	engine.AddFunc("live", Live)
	Engine = engine

}
