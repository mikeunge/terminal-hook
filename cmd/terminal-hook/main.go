package main

import (
	"os"

	"github.com/akamensky/argparse"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

const (
	AppName        = "terminal-hook"
	AppDescription = "Make cd great again."
	AppAuthor      = "@mikeunge"
	AppVersion     = "0.0.1"
	Github         = "https://github.com/mikeunge/terminal-hook"
)

func printAbout() {
	s, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("T", pterm.FgRed.ToStyle()),
		putils.LettersFromStringWithStyle("H", pterm.FgWhite.ToStyle())).
		Srender()
	pterm.DefaultCenter.Println(s)
	pterm.DefaultCenter.
		WithCenterEachLineSeparately().
		Printf("%s - v%s\n%s\n\n"+pterm.Red("Author:")+" %s\n"+pterm.Red("Repository:")+" %s\n", AppName, AppVersion, AppDescription, AppAuthor, Github)
}

func main() {
	parser := argparse.NewParser(AppName, AppDescription)
	argAbout := parser.Flag("", "about", &argparse.Options{Required: false, Help: "Print information about the app."})

	err := parser.Parse(os.Args)
	if err != nil {
		pterm.Error.Printf("Parsing error\n%+v", parser.Usage(err))
		os.Exit(1)
	}

	if *argAbout {
		printAbout()
		os.Exit(0)
	}

	pterm.Info.Println("Hello, World!")
}
