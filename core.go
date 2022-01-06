package GoLive

import (
	"github.com/gofiber/template/html"
	"log"
	"time"
)

type Map map[string]interface{}

// EventFunc type for event functions
type EventFunc func(state *State, data Map)

// Engine the engine used for rendering templates
var Engine *html.Engine

// Connections list of all connections
var Connections = make(map[string]*WebSocket)

// SetupEngine sets up the HTML engine for rendering templates
func SetupEngine(engine *html.Engine) {
	engine.AddFunc("live", Live)
	engine.AddFunc("live_child", LiveChild)
	Engine = engine
}

// SetupGarbageCollector initializes the garbage collector for removing old states
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
			}
			time.Sleep(10 * time.Second)
		}
	}()
}

// InitComponents loops over all components and calls the OnInit function
func InitComponents() {
	for _, component := range components {
		component.OnInit()
	}
}

// EmitEvent emits an event to all components.
func EmitEvent(name string, data Map) {
	for _, c := range components {
		c.Emit(name, data)
	}
}

// EmitEventTo emits an event to a specific component.
func EmitEventTo(componentName string, name string, data Map) {
	for _, c := range components {
		if c.GetName() == componentName {
			c.Emit(name, data)
		}
	}
}

// BroadcastEvent emits an event to all connections.
func BroadcastEvent(name string, data Map) {
	for _, c := range Connections {
		c.Send(Map{
			"event": name,
			"data":  data,
		})
	}
}

// BroadcastEventTo emits an event to a specific connection.
func BroadcastEventTo(SocketID string, name string, data Map) {
	for _, c := range Connections {
		if c.ID == SocketID && c.IsConnected {
			c.Send(Map{
				"event": name,
				"data":  data,
			})
		}
	}
}
