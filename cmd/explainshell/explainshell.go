package explainshell

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

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
		RequestBody := RequestBody{
			Model:            "text-davinci-003",
			Prompt:           "Explain this shell command in " + cmd.Flag("language").Value.String() + ": " + cmd.Flag("prompt").Value.String(),
			Temperature:      0,
			MaxTokens:        501,
			TopP:             1,
			FrequencyPenalty: 0,
			PresencePenalty:  0,
		}

		RequestBodyBytes, err := json.Marshal(RequestBody)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/davinci/completions", bytes.NewBuffer(RequestBodyBytes))
		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var TextCompletionResponse TextCompletionResponse
		err = json.Unmarshal(body, &TextCompletionResponse)
		if err != nil {
			panic(err)
		}

		if len(TextCompletionResponse.Choices) > 0 {
			println(TextCompletionResponse.Choices[0].Text)
		} else {
			println("No explanation found or OpenAI API key is invalid or missing.")
		}
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
