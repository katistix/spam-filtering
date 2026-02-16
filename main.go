package main

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func tokenize(message string) []string {
	tokens := strings.Fields(message)
	for i := range tokens {
		tokens[i] = strings.ToUpper(tokens[i])
	}
	return tokens
}

func addFileToBow(filePath string, bow map[string]int) error {
	// read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// increase the token appearances
	for _, token := range tokenize(string(content)) {
		bow[strings.ToUpper(token)] += 1
	}

	return nil
}

func main() {
	fmt.Printf("hello spam-filtering\n\n\n")

	// Add a dir to bow
	bow := map[string]int{}

	// add dir to bow
	err := filepath.WalkDir("./data/enron1", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		addFileToBow(path, bow)
		return nil
	})

	if err != nil {
		panic(err)
	}

	totalCount := 0
	for token := range bow {
		totalCount += bow[token]
	}

	// compute prob. of a document
	path := "./data/enron2/ham/0004.1999-12-10.kaminski.ham.txt"
	emailBow := map[string]int{}
	err = addFileToBow(path, emailBow)
	if err != nil {
		panic(err)
	}
	dp := 1.0
	for token := range emailBow {
		// if the current token is new, just ignore it
		if bow[token] == 0 {
			continue
		}

		p := float64(bow[token]) / float64(totalCount)
		dp += math.Log(p)
		fmt.Printf("%v => %v\n", token, p)
	}
	fmt.Printf("dp = %v\n", dp)

	// for token := range bow {
	// 	fmt.Printf("|%v| => %v\n", token, float64(bow[token])/float64(totalCount))
	// }
}
