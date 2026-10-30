package cmd

import (
	"github.com/shaninalex/repository-watcher/app/core"
	"github.com/spf13/cobra"
)

func NewRunWatcher(app *core.Instance) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run watcher",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
	}
	return cmd
}
