package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var exclude []string

var rootCmd = &cobra.Command{
	Use:               "go-ptool",
	Short:             "A collection of useful photo collection helpers",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	// Long: `A longer description that spans multiple lines`
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringArrayVarP(&exclude, "exclude", "x", []string{}, "patterns to exclude")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
