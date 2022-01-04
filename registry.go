package GoLive

import "log"

var components = make(map[string]IComponent)

func RegisterComponent(name string, component IComponent) {
	if _, ok := components[name]; ok {
		log.Printf("A component with name %s already exists\n", name)
		return
	}

	component.SetFile("components/" + name)
	component.SetName(name)
	components[name] = component
}

func GetComponent(name string) IComponent {
	if component, ok := components[name]; ok {
		return component
	}

	log.Printf("No component with name %s exists\n", name)
	return nil
}
