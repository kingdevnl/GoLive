package GoLive

import "time"

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
}

func (s *State) Set(name string, value interface{}) {
	s.data[name] = value
}
func (s State) Get(name string) interface{} {
	return s.data[name]
}

func NewState() *State {
	return &State{
		ID:        GenerateID(id_len),
		data:      make(Map),
		CreatedAt: time.Now(),
		LastUsed:  time.Now(),
	}
}
