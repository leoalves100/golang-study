package migrate

import (
	"api/src/config"
	"api/src/database/migrations"

	"github.com/spf13/cobra"
)

func Migration() *cobra.Command {
	var migrationCmd = &cobra.Command{
		Use:   "migration",
		Short: "Migration database",
		Run: func(cmd *cobra.Command, args []string) {
			config.LoadingConfigs()
			migrations.Migration()
		},
	}

	return migrationCmd
}
