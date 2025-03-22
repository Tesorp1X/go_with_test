package maps

import (
	"fmt"
)

func Search(dictionary map[string]string, word string) (string, error) {
	if val, ok := dictionary[word]; ok {
		return val, nil
	}
	return "", fmt.Errorf("error: couldn't find word %s", word)
}
