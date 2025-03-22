package maps

var (
	ErrNoSuchWord        = DictionaryError("error: no such word")
	ErrWordAlreadyExists = DictionaryError("error: word is already in the dictionary")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

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

func (d Dictionary) Update(word, newDefinition string) error {
	if _, ok := d[word]; ok {
		d[word] = newDefinition
		return nil
	}
	return ErrNoSuchWord
}

func (d Dictionary) Delete(word string) error {
	if _, ok := d[word]; ok {
		delete(d, word)
		return nil
	}
	return ErrNoSuchWord
}
