package cmd

import (
	"go.uber.org/zap"
	"os"

	"github.com/spf13/cobra"
)

var debug, displayClues, displayRowAndColNumbers bool

var puzzleFilePath string
var log *zap.SugaredLogger

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nonogo [command]",
	Short: "Nonogram solver",
	Long:  "A Nonogram solver written in Go. See https://github.com/paulwalk/nonogo for more information.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "", false, "--debug=true|false")
	rootCmd.PersistentFlags().BoolVarP(&displayClues, "displayClues", "", true, "--displayClues=true|false")
	rootCmd.PersistentFlags().BoolVarP(&displayRowAndColNumbers, "displayRowAndColNumbers", "", true, "--displayRowAndColNumbers=true|false")
	rootCmd.PersistentFlags().StringVarP(&puzzleFilePath, "puzzle", "", "", "--puzzle=path-to-puzzle-file")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(solveCmd)
}
