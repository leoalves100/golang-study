package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func GenBase64() *cobra.Command {
	var serverCmd = &cobra.Command{
		Use:   "base64",
		Short: "Generation base64 random",
		Run: func(cmd *cobra.Command, args []string) {
			key := make([]byte, 64)
			if _, err := rand.Read(key); err != nil {
				log.Fatal(err)
			}

			stringBase64 := base64.StdEncoding.EncodeToString(key)
			fmt.Println(stringBase64)
		},
	}

	return serverCmd
}
