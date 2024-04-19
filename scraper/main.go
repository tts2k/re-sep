package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func checkBoolFlagsConflict(flagList []string) error {
	hasFlagEnabled := false
	var enabledFlag string

	for _, flag := range flagList {
		value, _ := pflag.CommandLine.GetBool(flag)

		if hasFlagEnabled && value {
			return fmt.Errorf("conflicting flags: %s, %s", enabledFlag, flag)
		}

		if value {
			hasFlagEnabled = true
			enabledFlag = flag
		}
	}

	return nil
}

func initFlags() error {
	pflag.Usage = func() {
		fmt.Fprintln(os.Stderr,
			"CLI for the scraper of re-sep\n\n"+
				"Usage:\n"+
				"  re-sep-cli [flags] <url>\n\n"+
				"Flags:",
		)
		pflag.PrintDefaults()
	}

	pflag.BoolP("help", "h", false, "Print this help message")
	pflag.BoolP("all", "a", false, "Scrape all available articles")
	pflag.BoolP("single", "s", false, "Scrape a single article")
	pflag.CommandLine.SortFlags = false

	pflag.Parse()

	err := checkBoolFlagsConflict([]string{"json", "single"})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := initFlags()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		pflag.Usage()
		return
	}

	// Help or no flag
	helpF, _ := pflag.CommandLine.GetBool("help")
	if helpF || pflag.NFlag() == 0 {
		pflag.Usage()
		return
	}

	// Scrape all
	allF, _ := pflag.CommandLine.GetBool("all")
	if allF {
		fmt.Fprintln(os.Stderr, "Scrape all is not implemented")
		return
	}

	// Scrape once
	singleF, _ := pflag.CommandLine.GetBool("single")
	if !singleF {
		panic("no single flag detected when there should be one")
	}

	if pflag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Scrape all is not implemented")
		return
	}
}
