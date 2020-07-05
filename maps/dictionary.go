package main

import "errors"

// ErrNotFound stores an error for not found
var ErrNotFound = errors.New("could not find the word you were looking for")

// Dictionary is a custom type
type Dictionary map[string]string

// Search searches for a word in a map
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
