package jsongpt

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jsongpt",
	Short: "jsongpt - a simple CLI to mock JSON using GPT",
	Long: `jsongpt is a simple CLI to mock JSON using GPT.
	
	You can use GPT-3 or GPT-4 to generate JSON.
	Note that you have to provide your own API key.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
