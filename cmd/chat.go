package cmd

import (
	"log"

	"github.com/musl/chaat/lib/chaat"

	"github.com/spf13/cobra"
)

var (
	chatConfig = &chaat.ChatConfig{}

	chatCmd = &cobra.Command{
		Use:   `chat <url>`,
		Short: `Connect to a chaat server.`,
		Long:  `Connect to a chaat server.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			chatConfig.URL = args[0]
			log.Fatal(chaat.Client(chatConfig))
		},
	}
)

func init() {
	rootCmd.AddCommand(chatCmd)
}
