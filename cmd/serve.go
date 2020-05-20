package cmd

import (
	"log"

	"github.com/musl/chaat/lib/chaat"

	"github.com/spf13/cobra"
)

var (
	serveConfig = &chaat.ServeConfig{}

	serveCmd = &cobra.Command{
		Use:   `serve`,
		Short: `Run a chaat server.`,
		Long:  `Run a chaat server.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Fatal(chaat.Serve(serveConfig))
		},
	}
)

func init() {
	serveCmd.Flags().StringVarP(&(serveConfig.URL), "url", "u", "tcp://0.0.0.0:61381", "URL to bind to")
	rootCmd.AddCommand(serveCmd)
}
