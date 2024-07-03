package main

// Defining a custom type lets us use methods on it
type Dictionary map[string]string

// here's a "raw function" for Search
func Search(dictionary Dictionary, word string) string {
	return dictionary[word]
}

// here's a method for Search that can be used with type Dictionary
func (d Dictionary) Search(word string) string {
	return d[word]
}
