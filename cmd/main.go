package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var apiKey string
var proxyUrl string
var prompt string
var content string
var showResult int
var model string

func main() {
	rootCmd := &cobra.Command{
		Use:   "openai-cmd",
		Short: "openAi command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return requestOpenAi()
		},
	}
	rootCmd.PersistentFlags().StringVarP(&apiKey, "apiKey", "k", "", "-k key")
	rootCmd.PersistentFlags().StringVarP(&proxyUrl, "proxyUrl", "u", "", "-u sock5://127.0.0.1:7890")
	rootCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", "Please translate to zh-CN", "-p Please translate to")
	rootCmd.PersistentFlags().StringVarP(&content, "content", "c", "", "-c 你好")
	rootCmd.PersistentFlags().StringVarP(&model, "model", "m", "gpt-3.5-turbo", "-m gpt-3.5-turbo")
	rootCmd.PersistentFlags().IntVarP(&showResult, "showResult", "s", 0, "-s 0")
	err := rootCmd.MarkPersistentFlagRequired("apiKey")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	err = rootCmd.MarkPersistentFlagRequired("content")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	err = rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func requestOpenAi() error {
	var c *openai.Client

	pUrl, err := url.Parse(proxyUrl)
	if err == nil && len(proxyUrl) > 1 {
		config := openai.DefaultConfig(apiKey)
		transport := &http.Transport{Proxy: http.ProxyURL(pUrl)}
		config.HTTPClient = &http.Client{Transport: transport}
		c = openai.NewClientWithConfig(config)
	} else {
		c = openai.NewClient(apiKey)
	}

	resp, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: prompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
			MaxTokens:        3000,
			TopP:             1,
			FrequencyPenalty: 1,
			PresencePenalty:  1,
		},
	)

	if err != nil {
		return err
	}

	result := resp.Choices[0].Message.Content
	fmt.Println(result)
	if showResult == 1 {
		text := "Content: \n" + strings.Replace(content, "\"", "\\\"", -1) + "\n\n" + "Result: \n" + strings.Replace(result, "\"", "\\\"", -1)
		// fmt.Println(text)

		cmd := exec.Command("osascript", "-e", "display dialog \""+text+"\" with title \""+prompt+"\" buttons {\"Copy\", \"OK\"} default button \"OK\" cancel button \"Copy\"")
		output, _ := cmd.Output()
		// fmt.Println(output, err)
		if string(output) == "" || !strings.Contains(string(output), "OK") {
			clipboard.Write(clipboard.FmtText, []byte(result))
		}
	} else {
		clipboard.Write(clipboard.FmtText, []byte(result))
	}
	return err
}
