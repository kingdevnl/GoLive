package GoLive

import (
	"github.com/gofiber/template/html"
	"log"
	"time"
)

type Map map[string]interface{}
type EventFunc func(state *State, data Map)

var Engine *html.Engine
var Connections = make(map[string]*WebSocket)

func SetupEngine(engine *html.Engine) {
	engine.AddFunc("live", Live)
	engine.AddFunc("live_child", LiveChild)
	Engine = engine
}

func SetupGarbageCollector() {

	go func() {

		for {

			// loop over all states
			for id, state := range states {
				// check if the socket id is not set and the state is older than 5 minutes
				if state.SocketID == "" && time.Now().Sub(state.CreatedAt) > time.Minute*5 {
					log.Println("SetupGarbageCollector: Removing state: " + id)
					delete(states, id)
					continue
				}

				//// check if the state lastUse is older than 1 hour
				//if time.Now().Sub(state.LastUsed) > time.Hour {
				//	log.Println("SetupGarbageCollector: Removing state: " + id)
				//	delete(states, id)
				//	continue
				//}

			}
			time.Sleep(10 * time.Second)

		}
	}()
}

func InitComponents() {
	for _, component := range components {
		component.OnInit()
	}
}
