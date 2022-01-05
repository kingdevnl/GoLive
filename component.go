package GoLive

import (
	"html/template"
)

type EventMethod func(state *State, data Map)

type EventHandler struct {
	Name  string
	Func  EventMethod
	State *State
}

// IComponent Interface
type IComponent interface {
	// OnInit called once when component is registered.
	OnInit()

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

	// GetName returns the name of the component.
	GetName() string
	// SetName sets the name of the component.
	SetName(name string)

	// Register a method to the component. (Used for event handlers like button clicks)
	Register(name string, method EventMethod)
	// GetMethod gets a method from the component.
	GetMethod(name string) EventMethod

	// On will register an event handler to the component.
	On(state *State, name string, method EventMethod)
	// Emit will emit an event to the component.
	Emit(name string, data Map)

	// GetEvents Return all events that are registered to the component.
	GetEvents() []EventHandler

	// RemoveEventHandler removes an event handler from the component.
	RemoveEventHandler(index int)
}

// Component default component implementation.
type Component struct {
	file    string
	name    string
	methods map[string]EventMethod
	events  []EventHandler
}

// OnInit default component implementation.
func (c Component) OnInit() {}

// OnMount default component implementation.
func (c Component) OnMount(state *State, args []interface{}) {
}

// Render default component implementation.
func (c *Component) Render(state *State) template.HTML {
	return renderComponent(c, state)
}

// ReRender Re-Renders the component. and sends it to the client.
func (c Component) ReRender(state *State) {
	// get the connection from the state.
	socket := Connections[state.SocketID]

	// if the connection is available then render and send the updated html to the client.
	if socket != nil && socket.IsConnected {
		socket.Send(Map{
			"kind": "rerender",
			"id":   state.ID,
			"html": state.Component.Render(state),
		})
	}
}

// GetFile returns the file name of the component.
func (c Component) GetFile() string {
	return c.file
}

// SetFile sets the file name of the component.
func (c *Component) SetFile(s string) {
	c.file = s
}

// GetName returns the name of the component.
func (c Component) GetName() string {
	return c.name
}

// SetName sets the name of the component.
func (c *Component) SetName(s string) {
	c.name = s
}

// Register a method to the component. (Used for event handlers like button clicks)
func (c *Component) Register(name string, method EventMethod) {
	if c.methods == nil {
		c.methods = make(map[string]EventMethod)
	}
	c.methods[name] = method
}

// GetMethod gets a method from the component.
func (c Component) GetMethod(name string) EventMethod {
	if c.methods == nil {
		c.methods = make(map[string]EventMethod)
	}
	if method, ok := c.methods[name]; ok {
		return method
	}
	return nil
}

// On will register an event handler to the component.
func (c *Component) On(state *State, name string, method EventMethod) {
	if c.events == nil {
		c.events = make([]EventHandler, 0)
	}
	// add the event handler to the component.
	c.events = append(c.events, EventHandler{
		Name:  name,
		Func:  method,
		State: state,
	})
}

// Emit will emit an event to the component.
func (c *Component) Emit(name string, data Map) {
	if c.events == nil {
		c.events = make([]EventHandler, 0)
	}
	// loop through all the event listeners and emit the event.
	for _, event := range c.events {
		if event.Name == name {
			event.Func(event.State, data)
		}
	}
}

// GetEvents Return all events that are registered to the component.
func (c *Component) GetEvents() []EventHandler {
	return c.events
}

// RemoveEventHandler removes an event handler from the component.
func (c *Component) RemoveEventHandler(index int) {
	c.events = append(c.events[:index], c.events[index+1:]...)
}
