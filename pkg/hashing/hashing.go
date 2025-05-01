package hashing

import (
	"math"
)

const (
	base    = 62
	charSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// ToBase62 Hash generates a hash for the given input string.
func ToBase62(id int64) string {
	var shortURL []byte
	for id > 0 {
		shortURL = append(shortURL, charSet[id%base])
		id /= base
	}
	// reverse
	for i, j := 0, len(shortURL)-1; i < j; i, j = i+1, j-1 {
		shortURL[i], shortURL[j] = shortURL[j], shortURL[i]
	}
	return string(shortURL)
}

// FromBase62 decodes a base62 encoded string back to its original integer value.
func FromBase62(shortURL string) int64 {
	var id int64
	for i := 0; i < len(shortURL); i++ {
		idx := -1
		for j, c := range charSet {
			if c == rune(shortURL[i]) {
				idx = j
				break
			}
		}
		if idx == -1 {
			return -1
		}
		id = id*base + int64(idx)
	}
	return id
}

func GenerateShortURL(id int64) string {
	return ToBase62(id)
}

func GetMaxIDforDigits(digits int) int64 {
	return int64(math.Pow(base, float64(digits))) - 1
}
