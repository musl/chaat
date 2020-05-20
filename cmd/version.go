package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   `version`,
		Short: `Print version info.`,
		Long:  `Print version info.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println(Version, Revision)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
