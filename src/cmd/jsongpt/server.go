package jsongpt

import (
	"os"

	"github.com/itsmuntadhar/JsonGPT/server"
	"github.com/spf13/cobra"
)

var port string

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "Start the server",
	Args:    cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		server.Serve(port, apiKey)
	},
}

func init() {
	serverCmd.Flags().StringVarP(&apiKey, "apikey", "k", os.Getenv("OPENAI_API_KEY"), "API key for GPT-3 or GPT-4. Default is OPENAI_API_KEY environment variable")
	serverCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on. Default is 8080")
	rootCmd.AddCommand(serverCmd)
}
