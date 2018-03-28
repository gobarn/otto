package otto

// Scheduler type
type Scheduler struct{}

// NewScheduler creates a new scheduler
func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) start() {
}
