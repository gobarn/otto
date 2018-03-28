package otto

import (
	"fmt"
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
	config         Config
	started        bool
	workerManagers map[string]*WorkerManager
	scheduler      *Scheduler
}

// New create a new struct
func New() *Otto {
	return &Otto{
		workerManagers: make(map[string]*WorkerManager),
	}
}

// Start starts the otto broker
func (o *Otto) Start() {
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
}

// Enqueue a job
func (o *Otto) Enqueue(name string, job interface{}) {
	o.workerManagers[name] = NewWorkerManager()
}

func (o *Otto) startWorkerManagers() {
	for _, wm := range o.workerManagers {
		wm.start()
	}
}

// Stop will stop otto broker
func (o *Otto) Stop() {
}
