package GoLive

import (
	"html/template"
	"log"
)

type EventMethod func(state *State, data Map)

type EventHandler struct {
	Name  string
	Func  EventMethod
	State *State
}

// IComponent Interface
type IComponent interface {
	// OnMount is called the first time the component is rendered.
	OnMount(state *State, args []interface{})
	// Render returns the rendered component.
	Render(state *State) template.HTML

	// ReRender renders the component again. and sends it to the client
	ReRender(state *State)

	// GetFile returns the file name of the component.
	GetFile() string
	// SetFile sets the file name of the component.
	SetFile(string)

	GetName() string
	SetName(name string)

	Register(name string, method EventMethod)
	GetMethod(name string) EventMethod

	On(state *State, name string, method EventMethod)
	Emit(name string, data Map)

	GetEvents() []EventHandler
	RemoveEventHandler(index int)
}

// Component default component implementation.
type Component struct {
	file    string
	name    string
	methods map[string]EventMethod
	events  []EventHandler
}

func (c Component) OnMount(state *State, args []interface{}) {
}

func (c *Component) Render(state *State) template.HTML {
	return renderComponent(c, state)
}

func (c Component) ReRender(state *State) {

	log.Println("ReRender called", &state)
	socket := Connections[state.SocketID]
	if socket != nil && socket.IsConnected {
		socket.Send(Map{
			"kind": "rerender",
			"id":   state.ID,
			"html": state.Component.Render(state),
		})
	}
}

func (c Component) GetFile() string {
	return c.file
}

func (c *Component) SetFile(s string) {
	c.file = s
}
func (c Component) GetName() string {
	return c.name
}

func (c *Component) SetName(s string) {
	c.name = s
}

func (c *Component) Register(name string, method EventMethod) {
	if c.methods == nil {
		c.methods = make(map[string]EventMethod)
	}
	c.methods[name] = method
}

func (c Component) GetMethod(name string) EventMethod {
	if c.methods == nil {
		c.methods = make(map[string]EventMethod)
	}
	if method, ok := c.methods[name]; ok {
		return method
	}
	return nil
}

func (c *Component) On(state *State, name string, method EventMethod) {
	if c.events == nil {
		c.events = make([]EventHandler, 0)
	}
	c.events = append(c.events, EventHandler{
		Name:  name,
		Func:  method,
		State: state,
	})
}

func (c *Component) Emit(name string, data Map) {
	if c.events == nil {
		c.events = make([]EventHandler, 0)
	}
	for _, event := range c.events {
		if event.Name == name {
			event.Func(event.State, data)
		}
	}
}

func (c *Component) GetEvents() []EventHandler {
	return c.events
}
func (c *Component) RemoveEventHandler(index int) {
	c.events = append(c.events[:index], c.events[index+1:]...)
}
