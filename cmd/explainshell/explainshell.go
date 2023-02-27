package explainshell

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(explainCmd)
	explainCmd.Flags().String("prompt", "", "The command prompt to explain")
	explainCmd.MarkFlagRequired("prompt")
	explainCmd.Flags().String("language", "en", "The language of the command")
}

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Provides information about firewall rules for the environment",
	Long:  `Provides information about firewall rules for the environment`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Flag("prompt").Value.String())
		fmt.Println(cmd.Flag("language").Value.String())
	},
}

type RequestBody struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             float64 `json:"top_p"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
}

type TextCompletionResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   map[string]int
}

type Choice struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	LogProbs     interface{}
	FinishReason string `json:"finish_reason"`
}
