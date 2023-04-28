package processor

import (
	"github.com/sirupsen/logrus"
)

type Processor interface {
	Process(task Task)
}

type Task func()

// AsyncProcessor process tasks asynchronously
type AsyncProcessor struct {
	log               *logrus.Logger
	tasksBufferLength int
	goroutinesCount   int
	tasks             chan Task
}

// NewAsyncProcessor constructor
func NewAsyncProcessor(log *logrus.Logger, tasksBufferLength, goroutinesCount int) *AsyncProcessor {
	tasks := make(chan Task, tasksBufferLength)
	return &AsyncProcessor{
		log:               log,
		tasksBufferLength: tasksBufferLength,
		goroutinesCount:   goroutinesCount,
		tasks:             tasks,
	}
}

// Process add task to processing 
func (p *AsyncProcessor) Process(task Task) {
	p.log.Debugf("[%v] added to processing buffer", task)
	p.tasks <- task
}

// Start runs processing goroutines
func (p *AsyncProcessor) Start() {
	for i := 0; i < p.goroutinesCount; i++ {
		goroutineId := i
		go func(id int, tasks chan Task) {
			for task := range tasks {
				p.log.Debugf("[goroutine-%d] processing [%v]", goroutineId, task)
				task()
			}
		}(goroutineId, p.tasks)
	}
}
