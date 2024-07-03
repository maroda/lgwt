package main

// Defining a custom type lets us use methods on it
type Dictionary map[string]string

// type to implement the 'error' interface
type DictionaryErr string

// method implementing the 'error' interface on our DictionaryErr type
func (e DictionaryErr) Error() string {
	return string(e)
}

// fill DictionaryErr with error constants
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word already in dictionary")
	ErrWordDoesNotExist = DictionaryErr("cannot update, word does not exist")
)

// here's a "raw function" for Search
/*
func Search(dictionary Dictionary, word string) string {
	return dictionary[word]
}
*/

// here's a method for Search that can be used with type Dictionary
func (d Dictionary) Search(word string) (string, error) {
	// instead of assuming the index exists, we first actually check for it
	// this is the MAP LOOKUP:
	// 		If the key ('word') exists, ok will exist, definition == value
	// 		if the key does not, ok will not exist, defintion == ""
	//		'definition' is returned regardless, so we need to check for a key
	definition, ok := d[word]
	// This behaves like common error patterns:
	if !ok {
		return "", ErrNotFound
	}
	// return the value at the 'word' key, error == nil
	// return d[word], nil
	// return the value assigned to 'definition', error == nil
	return definition, nil
}

// new method: Add
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

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
