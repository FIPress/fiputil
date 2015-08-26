package fiputil

import "regexp"

var (
	emailR = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
)

func IsEmail(email string) bool {
	return emailR.MatchString(email)
}
