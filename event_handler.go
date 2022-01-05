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

	log.Println("Registered component:", componentName)
}

func _handleEvent(ws *WebSocket, data Map) {
	log.Println("handlEvent: ", data)

	//check if data contains event and type
	if data["event"] == nil || data["type"] == nil || data["id"] == nil || data["component"] == nil {
		return
	}

	event := data["event"].(string)
	eventType := data["type"].(string)

	componentId := data["id"].(string)
	componentName := data["component"].(string)

	component := GetComponent(componentName)
	state := states[componentId]

	if event == "register" {
		_handleRegister(ws, data)
		return
	}

	switch eventType {
	case "click":
		method := component.GetMethod(event)
		if method == nil {
			log.Println("Component method not found:", componentName, "#", event)
			return
		}
		method(state, data)
		return
	case "bind":
		state.Set(event, data["value"])
		component.ReRender(state)
	}

	//switch event {
	//case "register":
	//	_handleRegister(ws, data)
	//	break
	//}

}
