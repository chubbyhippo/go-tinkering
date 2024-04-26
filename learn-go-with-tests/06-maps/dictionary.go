package main

import (
	"errors"
)

type Dictionary map[string]string

var (
	ErrorNotFound       = errors.New("could not find the word you were looking for")
	ErrWordExists       = errors.New("cannot add word because it already exists")
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrorNotFound):
		d[word] = definition
	case err == nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrorNotFound):
		return ErrWordDoesNotExist
	case err == nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {

}
