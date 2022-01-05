package GoLive

import (
	"bytes"
	"golang.org/x/net/html"
	"html/template"
	"log"
	"strings"
)

// parse Parses the given html string and returns a html Node of the first element in body.
func _parse(input string) *html.Node {

	root, err := html.Parse(strings.NewReader(input))
	if err != nil {
		log.Println("Error parsing HTML:", err)
		return nil
	}
	var body *html.Node
	var f func(*html.Node)

	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "body" {
			body = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			f(child)
		}
	}
	f(root)

	return body

}

// _toHTML Returns a html string of the given html Node.
func _toHTML(node *html.Node) template.HTML {
	var b bytes.Buffer
	html.Render(&b, node)
	return template.HTML(b.String())
}

func _getChildren(node *html.Node) []*html.Node {
	var children []*html.Node

	var f func(*html.Node)
	f = func(node *html.Node) {
		children = append(children, node)
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			f(child)
		}
	}
	f(node)
	return children
}

// renderComponent Renders the given component and returns the rendered html.
func renderComponent(c IComponent, state *State) template.HTML {

	var data = map[string]interface{}{
		"c":     c,
		"state": state,
	}

	for k, v := range state.data {
		data[k] = v
	}

	var bffer bytes.Buffer
	if err := Engine.Render(&bffer, c.GetFile(), data); err != nil {
		log.Println("Error rendering template:", c.GetFile(), err)
		return template.HTML("error rendering template " + c.GetFile())
	}

	body := _parse(bffer.String())
	body.FirstChild.Attr = append(body.FirstChild.Attr, html.Attribute{Key: "live-id", Val: state.ID})
	body.FirstChild.Attr = append(body.FirstChild.Attr, html.Attribute{Key: "live-component", Val: c.GetName()})

	children := _getChildren(body)
	// loop through all children and check if they have the live-bind attribute

	for _, node := range children {
		if node.Type == html.ElementNode {
			for _, attr := range node.Attr {
				if attr.Key == "live-bind" {
					value := state.data[attr.Val]
					if value == nil {
						value = ""
					}
					node.Attr = append(node.Attr, html.Attribute{Key: "value", Val: value.(string)})
				}
			}

		}
	}

	return _toHTML(body)
}

func Live(name string, args ...interface{}) template.HTML {

	component := GetComponent(name)

	if component == nil {
		return template.HTML("Failed to find component: " + name)
	}

	state := NewState()
	states[state.ID] = state

	component.OnMount(state, args)

	return component.Render(state)
}

func LiveChild(name string, stateID string, root *State, args ...interface{}) template.HTML {
	component := GetComponent(name)
	if component == nil {
		return template.HTML("Failed to find component: " + name)
	}

	if root == nil {
		return template.HTML("Failed to find root state: " + stateID)
	}

	state := root.Children[stateID]

	if state == nil {
		state = NewState()
		state.Parent = root
		states[state.ID] = state
		root.Children[stateID] = state
		component.OnMount(state, args)
	}

	return component.Render(state)

}
