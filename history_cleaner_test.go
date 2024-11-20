package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"testing"
)

func createTestHistory() {
	_ = os.Remove("Test_history/zhistory")
	os.Link("Test_history/reset.txt", "Test_history/zhistory")
}

func filesAreDifferent(file1, file2 string) bool {
	f1, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	scanner1 := bufio.NewScanner(f1)
	scanner2 := bufio.NewScanner(f2)

	for scanner1.Scan() {
		scanner2.Scan()
		if !bytes.Equal(scanner1.Bytes(), scanner2.Bytes()) {
			return true
		}
	}
	return false
}

func TestMainFunc(t *testing.T) {
	createTestHistory()
	os.Args = append(os.Args, "--file=Test_history/zhistory")
	os.Args = append(os.Args, "--regexFile=regex_patterns.txt")
	// os.Args = append(os.Args, "--verbose=true")
	main()
	if filesAreDifferent("Test_history/zhistory", "Test_history/expected_results") {
		t.Error("files are not the same")
	}

	// cleanup
	os.Remove("Test_history/zhistory")
	os.Remove("Test_history/zhistory.bak")
}
