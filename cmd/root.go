package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "trash",
	Short: `trash is a CLI that prevents irreversible file deletions.
It acts by moving, rather than actually removing, files to a folder, which can be cleaned or recovered in the future.`,
	Version: "v0.0.1",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
