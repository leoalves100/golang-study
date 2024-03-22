package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "DevBook API",
}

func Execute() {

	rootCmd.AddCommand(Server(), GenBase64())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
