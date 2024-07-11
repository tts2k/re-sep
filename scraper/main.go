package main

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"re-sep-scraper/config"
	"re-sep-scraper/database"
	"re-sep-scraper/scraper"
	"re-sep-scraper/utils"

	"github.com/spf13/pflag"
)

func doSingle(url string) error {
	outputPath, _ := pflag.CommandLine.GetString("out")
	if outputPath == "" {
		return errors.New("flag: no output specified")
	}
	_, err := os.Stat(outputPath)
	if os.IsExist(err) || err == nil {
		ans := utils.PromptYN("File exists. Do you want to update it?")
		if !ans {
			return errors.New("Aborted")
		}
	}

	// Scrape
	fmt.Println("=> Scraping article")
	article, err := scraper.Single(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Single scrape failed")
		return err
	}

	// Create database
	fmt.Println("=> Creating database")
	database.InitDB(outputPath)
	err = database.CreateTable()
	if err != nil {
		return err
	}

	fmt.Println("=> Inserting row into the database")
	err = database.InsertArticle(article)
	if err != nil {
		return err
	}

	return nil
}

func doAll() error {
	outputPath, _ := pflag.CommandLine.GetString("out")
	if outputPath == "" {
		return errors.New("flag: no output specified")
	}
	_, err := os.Stat(outputPath)
	if os.IsExist(err) || err == nil {
		ans := utils.PromptYN("File exists. Do you want to update it?")
		if !ans {
			return errors.New("Aborted")
		}
	}

	wg, results, err := scraper.All()
	if err != nil {
		return err
	}

	database.InitDB(outputPath)
	fmt.Println("=> Creating database")
	err = database.CreateTable()
	if err != nil {
		return err
	}

	dbWg := &sync.WaitGroup{}
	dbWg.Add(1)
	go func() {
		for result := range results {
			fmt.Println("Inserting: ", result.Title)
			err := database.InsertArticle(result)
			if err != nil {
				fmt.Println(err)
			}
		}
		dbWg.Done()
	}()

	wg.Wait()
	close(results)
	dbWg.Wait()

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
	pflag.BoolVarP(&config.All, "all", "a", false, "Scrape all available articles")
	pflag.BoolVarP(&config.Single, "single", "s", false, "Scrape a single article")
	pflag.StringVarP(&config.Output, "out", "o", "", "Specify output path")
	pflag.BoolVar(&config.Yes, "yes", false, "Assume yes")
	pflag.BoolVarP(&config.Verbose, "verbose", "v", false, "Verbose output")
	pflag.IntVarP(&config.WorkerCount, "worker", "w", 4, "Number of scraping workers")
	pflag.Int64Var(&config.Sleep, "sleep", -1, "Adjust worker sleep time after each job")

	pflag.CommandLine.SortFlags = false

	// Parse
	pflag.Parse()

	// Check conflict
	if config.All && config.Single {
		return fmt.Errorf("cannot have all and single at the same time")
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
	if config.All {
		err = doAll()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		return
	}

	// Scrape once
	if !config.Single {
		fmt.Fprintln(os.Stderr, "No single flag detected when there should be one")
		return
	}

	if pflag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "An url must be provided for single mode")
		pflag.Usage()
		return
	}

	err = doSingle(pflag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
