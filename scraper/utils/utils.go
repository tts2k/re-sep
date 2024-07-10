package utils

import (
	"fmt"
	"strings"

	"re-sep-scraper/config"
)

func PromptYN(text string) bool {
	if config.Config.Yes {
		return true
	}

	var choice string
	fmt.Printf("%s [Y/N]:", text)
	fmt.Scanf("%s", &choice)

	return strings.ToLower(choice) == "y"
}
