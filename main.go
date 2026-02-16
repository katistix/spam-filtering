package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Bow map[string]int

func tokenize(message string) []string {
	tokens := strings.Fields(message)
	for i := range tokens {
		tokens[i] = strings.ToUpper(tokens[i])
	}
	return tokens
}

func addFileToBow(filePath string, bow Bow) error {
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

func addDirToBow(dirPath string, bow Bow) error {
	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		addFileToBow(path, bow)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func totalCount(bow Bow) int {
	count := 0

	for w := range bow {
		count += bow[w]
	}
	return count
}

func main() {

	// Add a dir to bow
	ham := Bow{}
	err := addDirToBow("./data/enron1/ham", ham)
	if err != nil {
		panic(err)
	}

	spam := Bow{}
	err = addDirToBow("./data/enron1/spam", spam)
	if err != nil {
		panic(err)
	}

	hamTotalCount := totalCount(ham)
	spamTotalCount := totalCount(spam)
	combinedTotalCount := hamTotalCount + spamTotalCount

	fmt.Printf("ham  = %v\n", float64(hamTotalCount)/float64(combinedTotalCount))
	fmt.Printf("spam = %v\n", float64(spamTotalCount)/float64(combinedTotalCount))

	// totalCount := 0
	// for token := range bow {
	// 	totalCount += bow[token]
	// }

	// // compute prob. of a document
	// path := "./data/enron2/ham/0004.1999-12-10.kaminski.ham.txt"
	// emailBow := Bow{}
	// err = addFileToBow(path, emailBow)
	// if err != nil {
	// 	panic(err)
	// }
	// dp := 1.0
	// for token := range emailBow {
	// 	// if the current token is new, just ignore it
	// 	if bow[token] == 0 {
	// 		continue
	// 	}

	// 	p := float64(bow[token]) / float64(totalCount)
	// 	dp += math.Log(p)
	// 	fmt.Printf("%v => %v\n", token, p)
	// }
	// fmt.Printf("dp = %v\n", dp)

	// for token := range bow {
	// 	fmt.Printf("|%v| => %v\n", token, float64(bow[token])/float64(totalCount))
	// }
}
