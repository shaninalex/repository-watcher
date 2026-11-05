package cmd

import (
	"github.com/shaninalex/repository-watcher/app/core"
	"github.com/spf13/cobra"
)

func NewSetup(app core.Resolve) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "Setup watcher",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
	}
	return cmd
}
