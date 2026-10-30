package bootstrap

import (
	"context"
	"time"

	"github.com/shaninalex/repository-watcher/app/cmd"
	"github.com/shaninalex/repository-watcher/app/core"
	"github.com/shaninalex/repository-watcher/app/platform/config"
	"github.com/spf13/cobra"
)

func NewCLI() (*cobra.Command, func()) {
	var app *core.Instance
	root := &cobra.Command{
		Use:           "repo_watcher",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			cfg, err := config.Load()
			if err != nil {
				return err
			}

			app, err = core.New(cmd.Context(), cfg)
			return err
		},
	}

	root.AddCommand(
		cmd.NewRunWatcher(app),
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
