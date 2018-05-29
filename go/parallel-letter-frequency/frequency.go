package letter

import (
	"sync"
)

// FreqMap letters counting map
type FreqMap map[rune]int

// Frequency counts letters in string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}

	return m
}

// ConcurrentFrequency counts letters in string concurrently
func ConcurrentFrequency(input []string) FreqMap {
	var wg sync.WaitGroup
	var cm sync.Map
	m := FreqMap{}

	for _, s := range input {

		wg.Add(1)

		go func(s string) {
			defer wg.Done()

			r := Frequency(s)

			for k, v := range r {
				vv, exist := cm.LoadOrStore(k, v)
				if exist {
					new := vv.(int) + v
					cm.Store(k, new)
				}
			}
		}(s)

		wg.Wait()
	}

	// wg.Wait()

	cm.Range(func(k, v interface{}) bool {
		m[k.(rune)] = v.(int)
		return true
	})

	return m
}
