package worker

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"time"
)

type Worker struct {
	jobs map[string]func(ctx context.Context) error
}

func New() *Worker {
	return &Worker{
		jobs: make(map[string]func(ctx context.Context) error),
	}
}

func (w *Worker) Periodic(ctx context.Context, jobName string, job func(ctx context.Context) error, interval time.Duration) error {
	if _, ok := w.jobs[jobName]; ok {
		return errors.Errorf("jobs with name %s alredy exits", jobName)
	}

	w.jobs[jobName] = job

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				delete(w.jobs, jobName)
				return
			case <-ticker.C:
				log.Printf("start job %s", jobName)
				if err := job(ctx); err != nil {
					log.Printf("job %s - err %v", jobName, err)
				}
				log.Printf("end job %s", jobName)
			}
		}
	}()

	return nil
}
