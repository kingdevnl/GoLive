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

// renderComponent Renders the given component and returns the rendered html.
func renderComponent(c Component, state *State) template.HTML {

	var bffer bytes.Buffer
	if err := Engine.Render(&bffer, c.GetFile(), state.data); err != nil {
		log.Println("Error rendering template:", c.GetFile(), err)
		return template.HTML("error rendering template " + c.GetFile())
	}

	body := _parse(bffer.String())
	body.FirstChild.Attr = append(body.FirstChild.Attr, html.Attribute{Key: "live-id", Val: state.ID})
	body.FirstChild.Attr = append(body.FirstChild.Attr, html.Attribute{Key: "live-component", Val: c.GetName()})
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
