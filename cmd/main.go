package main

import (
	"flag"
	"fmt"
	"frequencycounter/pkg/bytes"
	fc "frequencycounter/pkg/frequencycounter"
	"log"
	"os"
)

const outputFormat = "%d\t%s\n"

func main() {
	fileName := flag.String("file", "mobydick.txt", "specify file name")
	count := flag.Int("count", 20, "specify count of words")
	flag.Parse()

	file, err := os.ReadFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	source := bytes.SplitIntoWords(file)

	var counter fc.FrequencyCounter = fc.New(source)
	result := counter.MostFrequentN(*count)

	for _, w := range result {
		fmt.Printf(outputFormat, w.Count, w.Word)
	}
}
