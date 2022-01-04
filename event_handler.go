package GoLive

import (
	"log"
)

func _handleRegister(ws *WebSocket, data Map) {

	componentId := data["id"].(string)
	componentName := data["component"].(string)

	state := states[componentId]

	if state == nil {
		log.Println("State not found. requesting reload...")
		ws.Send(Map{
			"kind":   "action",
			"action": "reload",
		})
		return
	}

	if state.SocketID != "" {
		ws.Send(Map{
			"kind":   "action",
			"action": "reload",
		})
		return
	}
	state.SocketID = ws.ID
	state.Component = GetComponent(componentName)

}

func handleRegularEvent(ws *WebSocket, data Map) {

	componentId := data["id"].(string)
	componentName := data["component"].(string)
	state := states[componentId]
	component := GetComponent(componentName)

	if component == nil {
		log.Println("Component not found: " + componentName)
		return
	}
	if state == nil {
		return
	}

	component.OnEvent(data["type"].(string), data["name"].(string), data, state)

}

func _handleEvent(ws *WebSocket, data Map) {
	event := data["event"]
	log.Println("EVENT: " + event.(string))
	switch event {
	case "register":
		_handleRegister(ws, data)
		break

	case "event":
		handleRegularEvent(ws, data)
		break
	}

}
