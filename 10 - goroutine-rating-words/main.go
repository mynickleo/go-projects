package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

func countWords(someStr string, ch chan map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()

	wordCount := make(map[string]int)
	words := strings.Fields(someStr)

	for _, word := range words {
		wordCount[word]++
	}

	ch <- wordCount
}

func mergeMaps(dest, src map[string]int) {
	for word, count := range src {
		dest[word] += count
	}
}

func main() {
	filePath := "text_file.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	ch := make(chan map[string]int, 4)
	wordsMap := make(map[string]int)

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)

	textChunk := ""
	chunkSize := 1000

	for scanner.Scan() {
		textChunk += scanner.Text() + " "
		if len(strings.Fields(textChunk)) >= chunkSize {
			wg.Add(1)
			go countWords(textChunk, ch, &wg)
			textChunk = ""
		}
	}

	if len(textChunk) > 0 {
		wg.Add(1)
		go countWords(textChunk, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for partialMap := range ch {
		mergeMaps(wordsMap, partialMap)
	}

	type wordFreq struct {
		word  string
		count int
	}
	var wordList []wordFreq
	for word, count := range wordsMap {
		wordList = append(wordList, wordFreq{word, count})
	}
	sort.Slice(wordList, func(i, j int) bool {
		return wordList[i].count > wordList[j].count
	})

	fmt.Println("Word frequencies in descending order:")
	for _, wf := range wordList {
		fmt.Printf("%s: %d\n", wf.word, wf.count)
	}
}
