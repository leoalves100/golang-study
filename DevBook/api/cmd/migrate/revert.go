package migrate

import (
	"api/src/config"
	"api/src/database/migrations"

	"github.com/spf13/cobra"
)

func Revert() *cobra.Command {
	var revertCmd = &cobra.Command{
		Use:   "revert",
		Short: "Revert migrations",
		Run: func(cmd *cobra.Command, args []string) {
			config.LoadingConfigs()
			migrations.Revert()
		},
	}

	return revertCmd
}
