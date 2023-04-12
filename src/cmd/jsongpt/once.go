package jsongpt

import (
	"fmt"
	"os"

	lib "github.com/itsmuntadhar/JsonGPT/pkg"
	"github.com/spf13/cobra"
)

var apiKey string
var length int
var language string
var gptModel string

var onceCmd = &cobra.Command{
	Use:     "once",
	Aliases: []string{"o"},
	Short:   "Generate JSON once",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		u := lib.Request{
			APIKey:   apiKey,
			Language: &language,
			Length:   &length,
			GPTModel: &gptModel,
			Model:    args[0],
		}
		resp, err := lib.GetGPTResponse(u)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(resp)
	},
}

func init() {
	onceCmd.Flags().StringVarP(&apiKey, "apikey", "k", os.Getenv("OPENAI_API_KEY"), "API key for GPT-3 or GPT-4. Default is OPENAI_API_KEY environment variable")
	onceCmd.Flags().IntVarP(&length, "length", "l", 1, "Length of the generated JSON. Default is 1")
	onceCmd.Flags().StringVarP(&language, "language", "L", "english", "Language of the generated JSON. Default is english")
	onceCmd.Flags().StringVarP(&gptModel, "model", "m", "gpt-3.5-turbo", "GPT model to use. Default is gpt-3.5-turbo")
	rootCmd.AddCommand(onceCmd)
}
