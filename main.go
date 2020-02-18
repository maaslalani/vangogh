package main

import (
	"flag"
	"fmt"
	"os"

	"vangogh/generate"
	"vangogh/style"
	"vangogh/text"
)

func main() {
	textCommand := flag.NewFlagSet("text", flag.ExitOnError)
	styleCommand := flag.NewFlagSet("style", flag.ExitOnError)
	generateCommand := flag.NewFlagSet("generate", flag.ExitOnError)

	filepath := textCommand.String("file", "gopher.png", "Image file to turn into text.")

	content := styleCommand.String("content", "gopher.png", "Image to be stylized")
	painting := styleCommand.String("painting", "starry", "Name of the painting to be used as the stylistic transfer image")

	if len(os.Args) < 2 {
		fmt.Println("Incorrect usage of vangogh. Please use the help command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "text":
		textCommand.Parse(os.Args[2:])
	case "style":
		styleCommand.Parse(os.Args[2:])
	case "generate":
		generateCommand.Parse(os.Args[2:])
	case "help":
		fmt.Println()
		flag.PrintDefaults()
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	switch {
	case textCommand.Parsed():
		text.Command(*filepath)
	case styleCommand.Parsed():
		style.Command(*content, *painting)
	case generateCommand.Parsed():
		generate.Command()
	}
}
