package GoLive

import "time"

// states List of all states
var states = make(map[string]*State)

type State struct {
	// ID of the state
	ID string
	// ID of the socket assigned to the state
	SocketID string
	// Data associated with the state
	data Map

	// Component class of the state
	Component IComponent

	// Date the state was created (Used for cleaning up states that have not been used for a while)
	CreatedAt time.Time
	LastUsed  time.Time

	Parent   *State
	Children map[string]*State
}

// Set sets a variable in the state
func (s *State) Set(name string, value interface{}) {
	s.LastUsed = time.Now()
	s.data[name] = value
}

// Get gets a variable from the state
func (s State) Get(name string) interface{} {
	s.LastUsed = time.Now()
	return s.data[name]
}

// NewState creates a new state and assigns it an ID and returns it
func NewState() *State {
	return &State{
		ID:        GenerateID(id_len),
		data:      make(Map),
		CreatedAt: time.Now(),
		LastUsed:  time.Now(),
		Children:  make(map[string]*State),
	}
}
