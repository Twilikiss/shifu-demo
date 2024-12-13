// Package task
// @Author twilikiss 2024/12/13 17:19:19
package task

import (
	"github.com/go-co-op/gocron"
	"shifu-demo/log"
	"shifu-demo/logic"
	"time"
)

type Task struct {
	s *gocron.Scheduler
}

func NewTask() *Task {
	return &Task{
		s: gocron.NewScheduler(time.UTC),
	}
}

func (t *Task) Run(url string, minute int) {
	t.s.Every(minute).Minute().Do(func() {
		result := logic.NewMeasurement().GetMeasurement(url)
		log.Info("get measurement avg: ", result)
	})
}

func (t *Task) StartBlocking() {
	t.s.StartBlocking()
}

func (t *Task) Stop() {
	t.s.Stop()
}
