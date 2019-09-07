package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word becuase it already exists")
	ErrWordDoesNotExist = DictionaryErr("connot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

var dictionary = make(map[string]string)

func (d Dictionary) Search(word string) (string, error) {
	dictionary, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return dictionary, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}
