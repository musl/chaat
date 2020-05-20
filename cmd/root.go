package cmd

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	// AppName is the identifier used for configuration and log
	// messages.
	AppName = `chaat`

	// Version is the sematic version string of this app, set by the
	// build system.
	Version = ``

	// Revision is the id of the changeset from the version control
	// system, set by the build system.
	Revision = ``

	rootCmd = &cobra.Command{
		Use:   AppName,
		Short: `It's a simple chat app.`,
		Long:  `It's really an excuse to hack on something that uses ZeroMQ.`,
		Run: func(cmd *cobra.Command, args []string) {
			// The root command just shows help.
			cmd.Help()
		},
	}
)

func init() {
	// Add global/persistent flags, etc here.
}

// Execute runs the root command. It's the main entrypoint for the app
// and must execute or none of the subcommands will execute.
func Execute() error {
	rand.Seed(time.Now().UnixNano())

	log.SetOutput(os.Stderr)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix(AppName + ` `)

	return rootCmd.Execute()
}
