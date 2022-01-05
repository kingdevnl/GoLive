package GoLive

import "log"

// components Map of all registered components.
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

// GetComponent returns a component by name.
func GetComponent(name string) IComponent {
	if component, ok := components[name]; ok {
		return component
	}

	log.Printf("No component with name %s exists\n", name)
	return nil
}
