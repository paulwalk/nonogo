package cmd

import (
	"github.com/spf13/cobra"
)

var scratchCmd = &cobra.Command{
	Use: "scratch",
	Run: func(cmd *cobra.Command, args []string) {
		initialiseApplication()
		scratch()
	},
}

func scratch() {
	log.Info("Scratch starting...")

	log.Info("Scratch completed")
}
