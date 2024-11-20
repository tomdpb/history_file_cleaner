package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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

func checkLine(line string, regexes []regexp.Regexp, verbose bool) string {
	for _, re := range regexes {
		if re.MatchString(line) {
			if verbose {
				log.Println("Removed ", line)
			}
			return ""
		}
	}
	return line
}

func main() {
	filename := flag.String("file", "zhistory", "This is the file that will be modified.")
	regexFile := flag.String("regexFile", "regex_patterns.txt", "This is the file that contains the regular expressions to be used.")
	verbose := flag.Bool("verbose", false, "Setting to true will display what lines were removed.")
	flag.Parse()

	// create backup just in case
	backupFile := *filename + ".bak"
	os.Link(*filename, backupFile)
	os.Remove(*filename)
	history, err := os.Open(backupFile)
	if err != nil {
		log.Fatal(err)
	}
	defer history.Close()

	regexes := createRegexesFromFile(*regexFile)

	newHistory, err := os.Create(*filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(history)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		line := scanner.Text()
		line = checkLine(line, regexes, *verbose)
		if line == "" {
			continue
		} else {
			fmt.Fprintln(newHistory, line)
		}
	}
	newHistory.Close()

}
