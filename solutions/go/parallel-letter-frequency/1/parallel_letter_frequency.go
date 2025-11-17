package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a single text.
func Frequency(text string) FreqMap {
	freq := FreqMap{}
	for _, r := range text {
		freq[r]++
	}
	return freq
}

// ConcurrentFrequency counts frequency for multiple texts concurrently.
func ConcurrentFrequency(texts []string) FreqMap {
	if len(texts) == 0 {
		return FreqMap{}
	}

	result := FreqMap{}
	ch := make(chan FreqMap)
	var wg sync.WaitGroup

	// Launch a goroutine for each text
	for _, t := range texts {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			ch <- Frequency(text)
		}(t)
	}

	// Close channel once all goroutines finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Merge all partial maps into final result
	for m := range ch {
		for r, count := range m {
			result[r] += count
		}
	}

	return result
}
