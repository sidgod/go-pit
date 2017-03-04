package main

import (
	"os"
	"fmt"
	"bufio"
	"io/ioutil"
	"strings"
)

func three() {
	counts := make(map[string]int)
	newCounts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("Provide list of file to read from")
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error %v while reading file %s\n", err, file)
				continue
			}
			countLines(f, counts)
			f.Close()

			fileData, err := ioutil.ReadFile(file)
			if err == nil {
				for _, line := range strings.Split(string(fileData), "\n") {
					if len(line) > 0 {
						newCounts[line]++
					}
				}
			}
		}
	}
	printMap("Count Map", counts)
	printMap("New Count Map", newCounts)
}

func printMap(mapName string, countsMap map[string]int) {
	for k, v := range countsMap {
		fmt.Printf("%s %d\t%s\n", mapName, v, k)
	}
}

func countLines(f *os.File, counts map[string]int) {
	fmt.Println("Reading file " + f.Name())
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}
