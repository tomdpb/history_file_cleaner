package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func createRegexesFromFile(filename string) []regexp.Regexp {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer file.Close()

	regexes := []regexp.Regexp{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		regex := scanner.Text()
		regexes = append(regexes, *regexp.MustCompile(regex))
	}

	return regexes
}

func checkLine(line string, regexes []regexp.Regexp) string {
	for _, re := range regexes {
		if re.MatchString(line) {
			return ""
		}
	}
	return line
}

func main() {
	filename := os.Args[1]
	// create backup just in case
	backupFile := filename + ".bak"
	os.Link(filename, backupFile)
	os.Remove(filename)
	history, err := os.Open(backupFile)
	if err != nil {
		log.Fatal(err)
	}
	defer history.Close()

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	regexFilePath := filepath.Dir(ex) + "/regex_patterns.txt"
	regexes := createRegexesFromFile(regexFilePath)

	newHistory, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(history)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		line := scanner.Text()
		line = checkLine(line, regexes)
		if line == "" {
			continue
		} else {
			fmt.Fprintln(newHistory, line)
		}
	}
	newHistory.Close()

}
