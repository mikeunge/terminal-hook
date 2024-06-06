package main

import (
	"os"

	path "github.com/mikeunge/go/pkg/path-helper"
	hookStore "github.com/mikeunge/terminal-hook/pkg/store"

	"github.com/akamensky/argparse"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

const (
	AppName        = "terminal-hook"
	AppDescription = "Make cd great again."
	AppAuthor      = "@mikeunge"
	AppVersion     = "0.0.2"
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
	argHook := parser.String("", "hook", &argparse.Options{Required: false, Help: "Creates a new hook right where you are."})
	argDelete := parser.String("", "delete", &argparse.Options{Required: false, Help: "Delete a hook by its key."})

	if err := parser.Parse(os.Args); err != nil {
		pterm.Printf("Parsing error\n%+v", parser.Usage(err))
		os.Exit(1)
	}

	if *argAbout {
		printAbout()
		os.Exit(0)
	}

	storePath := path.SanitizePath("~/.local/share/terminal.hook")
	store, err := hookStore.New(storePath)
	if err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}

	if len(*argHook) > 0 {
		cwd, err := os.Getwd()
		if err != nil {
			pterm.Error.Println(err)
			os.Exit(8)
		}
		if err := store.WritePath(*argHook, cwd); err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}
		pterm.Info.Println("Hook created")
		os.Exit(0)
	}

	if len(*argDelete) > 0 {
		if err := store.DeletePath(*argDelete); err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}
		pterm.Info.Println("Deleted hook")
		os.Exit(0)
	}

	os.Exit(0)
}
