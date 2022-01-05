package main

import (
	"github.com/kingdevnl/GoLive"
	"log"
)

type MyComponent struct {
	GoLive.Component
}

func (c *MyComponent) OnMount(state *GoLive.State, args []interface{}) {
	log.Printf("OnMount\n")
	state.Set("counter", 0)

	state.Set("message", "This is a paragraph.")

	state.Set("search", "john@doe.com")

	c.Register("clicky", c.clicky)
	c.Register("test", c.test)

	c.On(state, "inc", func(s *GoLive.State, data GoLive.Map) {
		log.Println("on inc")

		s.Set("counter", s.Get("counter").(int)+1)

		c.ReRender(s)
	})
}

func (c *MyComponent) clicky(state *GoLive.State, data GoLive.Map) {
	log.Printf("clicky\n")

	c.Emit("inc", nil)

}

func (c *MyComponent) test(state *GoLive.State, data GoLive.Map) {
	state.Set("search", "this is a test")
	c.ReRender(state)
}

//func (c MyComponent) OnEvent(event string, name string, data GoLive.Map, state *GoLive.State) {
//	log.Println("OnEvent:", event, name)
//	if event == "click" && name == "clicky" {
//		counter := state.Get("counter").(int)
//		counter += 1
//		state.Set("counter", counter)
//		c.ReRender(state)
//	}
//	if event == "input" && name == "search" {
//		str := data["value"].(string)
//		state.Set("counter", len(str))
//		state.Set("search", str)
//		log.Println("SEARCH: ", str)
//		c.ReRender(state)
//
//	}
//}
