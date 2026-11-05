package runner

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/shaninalex/repository-watcher/app/core"
)

type Repository struct {
	Url      string
	Interval time.Duration
}

type Runner struct {
	Repositories []Repository
}

func (s *Runner) Start(ctx context.Context, app core.Resolve) {
	fmt.Println("Starting runner...")

	wg := sync.WaitGroup{}
	for _, repository := range s.Repositories {
		wg.Add(1)
		go func(repository Repository) {
			defer wg.Done()
			work(ctx, app, repository)
		}(repository)
	}

	wg.Wait()
}
