package otto

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Config for otto
type Config struct {
	RedisURL   string
	PoolSize   int
	NumWorkers int
	Interval   int
	Namespace  string
}

// Otto to hold data
type Otto struct {
	sync.Mutex
	config         *Config
	started        bool
	workerManagers map[string]*WorkerManager
	scheduler      *Scheduler
}

// New create a new struct
func New(config *Config) *Otto {
	return &Otto{
		config:         config,
		workerManagers: make(map[string]*WorkerManager),
	}
}

// Start starts the otto broker
func (o *Otto) Start() error {
	o.Lock()
	defer o.Unlock()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer close(signals)
		<-signals
		o.Stop()
	}()

	if !o.started {
		o.scheduler = NewScheduler()
		o.scheduler.start()

		o.startWorkerManagers()
		o.started = true
		fmt.Println("Otto started ...")
	}

	return nil
}

// Stop will stop otto broker
func (o *Otto) Stop() {
}

// Register a worker with broker
func (o *Otto) Register(worker *Worker) {
	log.Println("New worker registered.")
	o.workerManagers[worker.Queue] = NewWorkerManager(worker, o.config)
}

// Enqueue a job
func (o *Otto) Enqueue(name string, job Job) {
}

func (o *Otto) startWorkerManagers() {
	for _, wm := range o.workerManagers {
		wm.start()
	}
}
