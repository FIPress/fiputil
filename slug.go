package fiputil

import (
	"regexp"
	"strings"
)

var (
	invalid      = regexp.MustCompile(`[^\p{L}\p{N}]`) //\p{L} unicode letter, \p{N}: number
	multipleDash = regexp.MustCompile(`(\-){2,}`)
)

func Slug(title string) (slug string) {
	if len(title) == 0 {
		return
	}

	slug = strings.ToLower(title)
	slug = invalid.ReplaceAllString(slug, "-")
	slug = strings.Replace(slug, ".", "-", -1)
	slug = slugTrim(slug)

	slug = multipleDash.ReplaceAllString(slug, "-")

	return
}

func slugTrim(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		switch r {
		case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0, '-':
			return true
		}
		return false
	})
}
