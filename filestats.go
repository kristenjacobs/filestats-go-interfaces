package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type Stat interface {
	nextLine(line *string)
	printStats()
}

type LineCount struct {
	lineCount int
}

func (s *LineCount) nextLine(line *string) {
	s.lineCount++
}

func (s LineCount) printStats() {
	fmt.Printf("The line count is: %d\n", s.lineCount)
}

type WordCount struct {
	wordCount int
}

func (s *WordCount) nextLine(line *string) {
	s.wordCount += len(strings.Fields(*line))
}

func (s WordCount) printStats() {
	fmt.Printf("The word count is: %d\n", s.wordCount)
}

type AverageLettersPerWord struct {
	numLetters int
	numWords   int
}

func (s *AverageLettersPerWord) nextLine(line *string) {
	for _, char := range *line {
		if unicode.IsLetter(char) {
			s.numLetters++
		}
	}
	s.numWords += len(strings.Fields(*line))
}

func (s AverageLettersPerWord) printStats() {
	var alpw = 0.0
	if s.numWords > 0 {
		alpw = float64(s.numLetters) / float64(s.numWords)
	}
	fmt.Printf("The average number of letters per word is: %.2f\n", alpw)
}

type MostCommonLetter struct {
	letterFrequencyMap map[rune]int
}

func (s *MostCommonLetter) nextLine(line *string) {
	for _, char := range *line {
		if unicode.IsLetter(char) {
			s.letterFrequencyMap[unicode.ToLower(char)]++
		}
	}
}

func (s MostCommonLetter) printStats() {
	var maxVal = -1
	var maxKey = 'a'
	for key, val := range s.letterFrequencyMap {
		if val > maxVal {
			maxKey = key
			maxVal = val
		}
	}
	if maxVal != -1 {
		fmt.Printf("Most common letter is: %c\n", maxKey)
	} else {
		fmt.Printf("Could not calculate the most common letter\n")
	}
}

func openFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func processFile(fileName string, stats []Stat) {
	file := openFile(fileName)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		for _, stat := range stats {
			stat.nextLine(&line)
		}
	}
	file.Close()
	for _, stat := range stats {
		stat.printStats()
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf("Error: Invalid arguments.\n")
		fmt.Printf("  Usage:\n")
		fmt.Printf("    filestats-go-interfaces <file>\n")
		os.Exit(1)
	}
	fileName := args[0]

	processFile(fileName, []Stat{
		&LineCount{},
		&WordCount{},
		&AverageLettersPerWord{},
		&MostCommonLetter{letterFrequencyMap: make(map[rune]int)}})
}
