package otto

import (
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"
)

type Config struct {
    RedisURL   string
    PoolSize   int
    NumWorkers int
    Interval   int
    Namespace  string
}

type Otto struct {
    sync.Mutex
    config  Config
    started bool
}

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

        o.workerManager = NewWorkerManager()
        o.workerManager.start()

        o.started = true
    }
}

func (o *Otto) Stop() {
}
