package runner

import (
	"context"
	"fmt"
	"time"

	"github.com/shaninalex/repository-watcher/app/core"
)

func work(ctx context.Context, app core.Resolve, r Repository) {
	// something like:
	// events.RegisterProcess(ctx, app.Worker)

	ticker := time.NewTicker(r.Interval)
	for {
		select {
		case <-ticker.C:
			process(r)
		case <-ctx.Done():
			return
		}
	}
}

func process(r Repository) {
	fmt.Println("Processing repository: ", r.Url)
	// notify.Push(ctx, data)
}
