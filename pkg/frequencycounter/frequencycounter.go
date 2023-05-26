package frequencycounter

import (
	"bytes"
	"sort"
)

var _ FrequencyCounter = (*frequencyCounter)(nil)

// FrequencyCounter used to count words frequency
type FrequencyCounter interface {
	// MostFrequentN returns the N most frequent words in the source of frequencyCounter
	MostFrequentN(int) []WordFrequency
}

// frequencyCounter implements FrequencyCounter interface
type frequencyCounter struct {
	source          [][]byte
	wordFrequencies []WordFrequency
}

// WordFrequency stores the word as slice of bytes and it's count
type WordFrequency struct {
	Word  []byte
	Count uint64
}

// New creates frequencyCounter structure wit given source words
func New(source [][]byte) *frequencyCounter {
	return &frequencyCounter{
		source: source,
	}
}

func (c *frequencyCounter) MostFrequentN(n int) []WordFrequency {
	c.lowerSource()
	c.sortSource()
	c.countWordsFrequencies()
	c.sortWordsFrequencies()

	result := make([]WordFrequency, 0, n)
	for i := 0; i < n && i < len(c.wordFrequencies); i++ {
		result = append(result, c.wordFrequencies[i])
	}
	return result
}

// countWordsFrequencies counts the number of sequential repeating words in array
func (c *frequencyCounter) countWordsFrequencies() {
	var count uint64 = 1
	for i := 0; i < len(c.source); i++ {
		// append the last element, cause there are no next element to compare
		if i == len(c.source)-1 {
			c.wordFrequencies = append(c.wordFrequencies, WordFrequency{c.source[i], count})
			break
		}
		if bytes.Equal(c.source[i], c.source[i+1]) {
			count++
			continue
		}
		c.wordFrequencies = append(c.wordFrequencies, WordFrequency{c.source[i], count})
		count = 1
	}
}

func (c *frequencyCounter) sortSource() {
	sort.Slice(c.source, func(i, j int) bool {
		return bytes.Compare(c.source[i], c.source[j]) == 1
	})
}

func (c *frequencyCounter) sortWordsFrequencies() {
	sort.Slice(c.wordFrequencies, func(i, j int) bool {
		return c.wordFrequencies[i].Count > c.wordFrequencies[j].Count
	})
}

func (c *frequencyCounter) lowerSource() {
	for i := 0; i < len(c.source); i++ {
		c.source[i] = bytes.ToLower(c.source[i])
	}
}
