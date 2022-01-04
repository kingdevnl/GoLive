package main

import (
	"github.com/kingdevnl/GoLive"
	"log"
)

type MyComponent struct {
	GoLive.Component
}

func (c MyComponent) OnMount(state *GoLive.State, args []interface{}) {
	log.Printf("OnMount\n")
	state.Set("counter", 0)
	state.Set("message", "This is a paragraph.")
}

func (c MyComponent) OnEvent(event string, name string, data GoLive.Map, state *GoLive.State) {
	log.Println("OnEvent:", event, name)
	if event == "click" && name == "clicky" {
		counter := state.Get("counter").(int)
		counter += 1
		state.Set("counter", counter)
		c.ReRender(state)
	}
	if event == "bind" && name == "search" {
		str := data["value"].(string)
		state.Set("counter", len(str))
		state.Set("search", str)
		log.Println("SEARCH: ", str)
		c.ReRender(state)

	}
}
