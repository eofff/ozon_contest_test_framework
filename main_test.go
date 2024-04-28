package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

type TestData struct {
	FileIn  string
	FileA   string
	FileOut string
}

func Step(t *testing.T, testData TestData) {
	fileIn, _ := os.Open(testData.FileIn)
	defer fileIn.Close()
	if _, err := os.Stat(testData.FileOut); err == nil {
		os.Remove(testData.FileOut)
	}
	fileOut, _ := os.Create(testData.FileOut)
	defer fileOut.Close()

	Magic(fileIn, fileOut)

	answers, _ := os.ReadFile(testData.FileOut)
	expected, _ := os.ReadFile(testData.FileA)

	answersString := strings.TrimSpace(string(answers))
	expectedString := strings.TrimSpace(string(expected))

	if answersString != expectedString {
		t.Errorf("Error in %s test", testData.FileIn)
	}
}

func TestManual(t *testing.T) {
	var testData TestData
	testData.FileIn = "test.in"
	testData.FileOut = "test.out"
	testData.FileA = "test.a"
	Step(t, testData)
}

func TestFolder(t *testing.T) {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	dir := ""
	for _, e := range entries {
		if e.IsDir() {
			dir = e.Name()
		}
	}

	entries, err = os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	files := make([]string, 0)
	for _, e := range entries {
		if !e.IsDir() {
			name := e.Name()
			if len(name) < 2 || (name[len(name)-2:] != ".a" && name[len(name)-2:] != ".o") {
				files = append(files, e.Name())
			}
		}
	}

	for _, filename := range files {
		var testData TestData
		testData.FileIn = dir + "/" + filename
		testData.FileA = dir + "/" + filename + ".a"
		testData.FileOut = dir + "/" + filename + ".o"
		Step(t, testData)
	}
}
