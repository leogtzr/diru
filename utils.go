package main

import (
	"unicode/utf8"

	"github.com/spf13/viper"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)

	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}

	return string(buf)
}

func idToShortURL(id int, mChars []rune) string {
	shortURL := ""
	mapCharsSize := len(mChars)

	for id > 0 {
		shortURL += string(mChars[id%mapCharsSize])
		id /= mapCharsSize
	}

	return reverse(shortURL)
}

func shortURLToID(shortURL string, mChars []rune) int {
	mapCharsSize := len(mChars)
	id := 0

	for _, char := range shortURL {
		switch {
		case char >= 'a' && char <= 'z':
			id = id*mapCharsSize + int(char-'a')
		case char >= 'A' && char <= 'Z':
			id = id*mapCharsSize + int(char-'A') + 26
		case char >= '0' && char <= '9':
			id = id*mapCharsSize + int(char-'0') + 52
		}
	}

	return id
}

func readConfig(filename, configPath string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()

	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.SetConfigName(filename)
	v.AddConfigPath(configPath)
	v.SetConfigType("env")

	err := v.ReadInConfig()

	return v, err
}

func urlsToFullStat(urls *[]URLStat) []URLStatFull {
	urlFull := make([]URLStatFull, 0)

	for _, u := range *urls {
		shortURL := idToShortURL(u.ShortID, chars)
		urlFull = append(urlFull, URLStatFull{
			ShortURL:    shortURL,
			OriginalURL: u.URL,
		})
	}

	return urlFull
}
