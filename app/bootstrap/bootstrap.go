package bootstrap

import (
	"context"
	"time"

	"github.com/shaninalex/repository-watcher/app/cmd"
	"github.com/shaninalex/repository-watcher/app/core"
	"github.com/shaninalex/repository-watcher/app/platform/config"
	"github.com/spf13/cobra"
)

type Instance struct {
	Core *core.App
}

func (s *Instance) Close(ctx context.Context) error {
	return nil
}

func New(ctx context.Context, cfg *config.Config) (*Instance, error) {
	return &Instance{}, nil
}

func NewCLI() (*cobra.Command, func()) {
	var app *Instance
	root := &cobra.Command{
		Use:           "repo_watcher",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			cfg, err := config.Load()
			if err != nil {
				return err
			}

			app, err = New(cmd.Context(), cfg)
			return err
		},
	}

	resolve := core.Resolve(func() *core.App { return app.Core })
	root.AddCommand(
		cmd.NewRunWatcher(resolve),
		cmd.NewSetup(resolve),
	)

	cleanup := func() {
		if app != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			_ = app.Close(ctx)
		}
	}
	return root, cleanup
}
