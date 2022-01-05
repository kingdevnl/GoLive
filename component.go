package GoLive

import (
	"html/template"
	"log"
)

type Method func(state *State, data Map)

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

	Register(name string, method Method)
	GetMethod(name string) Method
}

// Component default component implementation.
type Component struct {
	file    string
	name    string
	methods map[string]Method
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

func (c *Component) Register(name string, method Method) {
	if c.methods == nil {
		c.methods = make(map[string]Method)
	}
	c.methods[name] = method
}

func (c Component) GetMethod(name string) Method {
	if c.methods == nil {
		c.methods = make(map[string]Method)
	}
	if method, ok := c.methods[name]; ok {
		return method
	}
	return nil
}
