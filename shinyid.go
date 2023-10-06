package shinyid

import (
	"errors"
)

const (
	base64URLCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
)

var (
	base64URLCharsetIndex [256]byte               // Lookup table for character indices
	base64URLCharsetLen   = len(base64URLCharset) // Length of the character set
)

func init() {
	for i := 0; i < base64URLCharsetLen; i++ {
		base64URLCharsetIndex[base64URLCharset[i]] = byte(i)
	}
}

// ToShiny converts an id to a shiny.
func ToShiny(id uint64) string {

	if id == 0 {
		return "A"
	}

	var shiny [11]byte // Fixed-length array to store the shiny
	i := 10            // Index for storing characters in the shiny array

	for id > 0 {
		shiny[i] = base64URLCharset[id&0x3F] // Extract the character from the character set
		id >>= 6                             // Shift the id by 6 bits to the right
		i--                                  // Decrement the index
	}

	return string(shiny[i+1:]) // Convert the shiny array to a string
}

// ToId converts a shiny to its corresponding id.
func ToId(shiny string) (uint64, error) {

	if shiny == "A" {
		return 0, nil
	}

	// Check if the shiny is a valid shiny
	if !isValidShiny(shiny) {
		return 0, errors.New("input must be a valid shiny")
	}

	// Convert the shiny to base10 id
	base10 := uint64(0)
	for i := 0; i < len(shiny); i++ {
		base64 := shiny[i]
		base64Value := base64URLCharsetIndex[base64]
		base10 = (base10 << 6) | uint64(base64Value)
	}

	return base10, nil
}

// isValidShiny checks if a given shiny is a valid shiny.
func isValidShiny(shiny string) bool {

	for _, char := range shiny {
		if !isAlphaNumeric(char) && char != '-' && char != '_' {
			return false
		}
	}
	return true
}

// isAlphaNumeric checks if a character is an alphanumeric character.
func isAlphaNumeric(char rune) bool {
	return ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') || ('0' <= char && char <= '9')
}
