package cmd

import (
	"time"

	"github.com/shaninalex/repository-watcher/app/core"
	"github.com/shaninalex/repository-watcher/app/core/runner"
	"github.com/spf13/cobra"
)

func NewRunWatcher(app core.Resolve) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run watcher",
		RunE: func(cmd *cobra.Command, _ []string) error {
			runner := &runner.Runner{
				Repositories: []runner.Repository{
					{
						Url:      "https://github.com/shaninalex/repository-watcher.git",
						Interval: time.Second * 2,
					},
					{
						Url:      "git@github.com:shaninalex/Lumna.git",
						Interval: time.Second * 5,
					},
				},
			}
			runner.Start(cmd.Context(), app)

			return nil
		},
	}
	return cmd
}
