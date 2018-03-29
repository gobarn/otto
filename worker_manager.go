package otto

// WorkerManager type
type WorkerManager struct {
	config *Config
}

// NewWorkerManager creates a new worker manager
func NewWorkerManager(worker *Worker, config *Config) *WorkerManager {
	return &WorkerManager{
		config: config,
	}
}

func (wm *WorkerManager) start() {
}
