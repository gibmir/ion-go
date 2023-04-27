package processor

import (
	"github.com/sirupsen/logrus"
)

type Processor interface {
	Process(task Task)
}

type Task func()

type AsyncProcessor struct {
	log               *logrus.Logger
	tasksBufferLength int
	goroutinesCount   int
	tasks             chan Task
}

func NewAsyncProcessor(log *logrus.Logger, tasksBufferLength, goroutinesCount int) *AsyncProcessor {
	tasks := make(chan Task, tasksBufferLength)
	return &AsyncProcessor{
		log:               log,
		tasksBufferLength: tasksBufferLength,
		goroutinesCount:   goroutinesCount,
		tasks:             tasks,
	}
}

func (p *AsyncProcessor) Process(task Task) {
	logrus.Debugf("[%v] added to processing buffer", task)
	p.tasks <- task
}

func (p *AsyncProcessor) Start() {
	for i := 0; i < p.goroutinesCount; i++ {
		goroutineId := i
		go func(id int, tasks chan Task) {
			for task := range tasks {
				logrus.Debugf("[goroutine-%d] processing [%v]", goroutineId, task)
				task()
			}
		}(goroutineId, p.tasks)
	}
}
