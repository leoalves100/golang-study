package cmd

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func Server() *cobra.Command {
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start the http server",
		Run: func(cmd *cobra.Command, args []string) {
			config.LoadingConfigs()
			r := router.RouterConfig()

			fmt.Printf("Rodando servidor na porta: %v\n", config.ExecutionPortAPI)
			log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.ExecutionPortAPI), r))
		},
	}

	return serverCmd
}
