package otto

type Handler func(m *Message)

// Worker type
type Worker struct {
	Queue       string
	Handler     Handler
	Concurrency int
}
