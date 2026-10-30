package core

import (
	"context"

	"github.com/shaninalex/repository-watcher/app/platform/config"
)

type Instance struct{}

func (s *Instance) Close(ctx context.Context) error {
	return nil
}

func New(ctx context.Context, cfg *config.Config) (*Instance, error) {
	return &Instance{}, nil
}
