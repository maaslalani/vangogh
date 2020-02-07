package main

import (
  "os"
  "fmt"
  "flag"

  "vangogh/text"
)

func main() {
  textCommand := flag.NewFlagSet("text", flag.ExitOnError)
  styleCommand := flag.NewFlagSet("style", flag.ExitOnError)

  filepath := textCommand.String("file", "gopher.png", "Image file to turn into text.")

  if len(os.Args) < 2 {
    fmt.Println("Incorrect usage of vangogh. Please use the help command")
    os.Exit(1)
  }

  switch os.Args[1] {
  case "text":
    textCommand.Parse(os.Args[2:])
  case "style":
    styleCommand.Parse(os.Args[2:])
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
    fmt.Println("Called style")
  }
}
