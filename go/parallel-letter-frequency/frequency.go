package letter

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
	c := make(chan FreqMap, len(input))
	m := FreqMap{}

	for _, s := range input {

		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	for i := 0; i < len(input); i++ {
		for k, v := range <-c {
			m[k] += v
		}
	}

	return m
}
