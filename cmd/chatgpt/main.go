package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/axetroy/openai"
	"github.com/gookit/color"
	"github.com/pkg/errors"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func checkExitCommand(question string) string {
	quest := strings.Trim(question, " ")
	keywords := []string{"", "loop", "break", "continue", "cls", "exit", "block"}
	for _, x := range keywords {
		if quest == x {
			return ""
		}
	}
	return quest
}

func printHelp() {
	println(`chatgpt - a cli for use chatGPT

USAGE:
chatgpt [OPTIONS]

OPTIONS:
  --help                 Print help information
  --version              Print version information
  --no-color             Disabled color for printing
  --model=[value]        The model param for chatGPT, defaults: text-davinci-003

SOURCE CODE:
  https://github.com/axetroy/chatGPT-cli`)
}

func run() error {
	var (
		showHelp    bool
		showVersion bool
		noColor     bool
		model       string
	)

	flag.BoolVar(&noColor, "no-color", false, "disabled color for printing")
	flag.BoolVar(&showHelp, "help", false, "Print help information")
	flag.BoolVar(&showVersion, "version", false, "Print version information")
	flag.StringVar(&model, "model", "text-davinci-003", "")

	flag.Usage = printHelp

	flag.Parse()

	if showHelp {
		printHelp()
		os.Exit(0)
	}

	if showVersion {
		println(fmt.Sprintf("%s %s %s", version, commit, date))
		os.Exit(0)
	}

	if color.SupportColor() {
		color.Enable = !noColor
	} else {
		color.Enable = false
	}

	apiKey := os.Getenv("OPEN_AI_API_KEY")

	if apiKey == "" {
		fmt.Println("chatgpt require OPEN_AI_API_KEY environment variable.")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	quit := false

	for !quit {
		fmt.Print("请输入你的问题(CTRL+C 退出): ")

		if !scanner.Scan() {
			break
		}

		question := scanner.Text()
		questionParam := checkExitCommand(question)
		switch questionParam {
		case "quit":
			quit = true
		case "":
			continue
		default:
			client := openai.NewClient(apiKey)

			err := client.CompletionsStream(openai.CompletionParams{
				Prompt:      question,
				Model:       model,
				Temperature: 0.6,
				MaxTokens:   2048,
			}, os.Stdout)

			if err != nil {
				return errors.WithStack(err)
			}
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(255)
	}
}
