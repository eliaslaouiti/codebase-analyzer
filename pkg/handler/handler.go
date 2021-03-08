package handler

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func AnalyzeCode() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file : %v", err)
	}

	repoURL := os.Getenv("REPOSITORY_URL")
	if repoURL == "" {
		log.Fatal("Missing required environment variable REPOSITORY_URL")
	}

	repoBranch := os.Getenv("REPOSITORY_BRANCH")
	if repoBranch == "" {
		repoBranch = "master"
	}

	toExclude := []string{}
	if os.Getenv("FILES_TO_EXCLUDE") != "" {
		toExclude = strings.Split(os.Getenv("FILES_TO_EXCLUDE"), ",")
	}

	dir := os.Getenv("DIRECTORY")
	if dir == "" {
		dir = "."
	}

	results, err := ExtractFlagsInfo(dir, toExclude)

	if err != nil {
		log.Fatalf("Error occured when parsing code files: %v", err)
	}

	for _, r := range results {
		log.Printf("Scanned file %s and found %d flag usages", r.File, len(r.Results))
	}

}
