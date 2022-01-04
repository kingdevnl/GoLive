package GoLive

type IWebSocket interface {
	IsConnected() bool
	ID() string
	Send(message interface{}) error
	Receive() (interface{}, error)
	Close() error
}
