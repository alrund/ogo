package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type Word struct {
	Value string
	Num   int
}

var reg = regexp.MustCompile(`[^а-я-]`)

func Top10(text string) []string {
	text = clean(text)
	counters := calculate(text)
	words := filter(counters)

	if len(words) == 0 {
		return nil
	}

	sort.Slice(words, func(i, j int) bool {
		switch {
		case words[i].Num > words[j].Num:
			return true
		case words[i].Num == words[j].Num:
			return words[i].Value < words[j].Value
		default:
			return false
		}
	})

	result := make([]string, 0, 10)
	for _, word := range words[:10] {
		result = append(result, word.Value)
	}

	return result
}

func clean(text string) string {
	text = strings.ToLower(text)
	text = reg.ReplaceAllString(text, " ")
	return text
}

func calculate(text string) map[string]*Word {
	strs := strings.Split(text, " ")
	counters := make(map[string]*Word, len(strs))
	for _, str := range strs {
		if _, ok := counters[str]; !ok {
			counters[str] = &Word{Value: str}
		}
		counters[str].Num++
	}
	return counters
}

func filter(counters map[string]*Word) []*Word {
	words := make([]*Word, 0, len(counters))
	for _, counter := range counters {
		if counter.Value == "" || counter.Value == "-" {
			continue
		}
		words = append(words, counter)
	}

	return words
}
