package maps

import (
	"errors"
)

var (
	ErrNoSuchWord = errors.New("error: no such word")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	if val, ok := d[word]; ok {
		return val, nil
	}
	return "", ErrNoSuchWord
}
