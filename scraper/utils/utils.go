package utils

import (
	"fmt"
	"strings"

	"re-sep-scraper/config"
)

func PromptYN(text string) bool {
	if config.Yes {
		return true
	}

	var choice string
	fmt.Printf("%s [Y/N]: ", text)
	_, err := fmt.Scanf("%s", &choice)
	if err != nil {
		panic(err)
	}

	return strings.ToLower(choice) == "y"
}
