package maps

import (
	"errors"
)

var (
	ErrNoSuchWord        = errors.New("error: no such word")
	ErrWordAlreadyExists = errors.New("error: word is already in the dictionary")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	if val, ok := d[word]; ok {
		return val, nil
	}
	return "", ErrNoSuchWord
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNoSuchWord:
		d[key] = value
	case nil:
		return ErrWordAlreadyExists
	default:
		return err
	}
	return nil
}
