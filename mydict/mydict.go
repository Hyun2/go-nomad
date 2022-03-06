package mydict

import "errors"

type Dictionary map[string]string

var (
	errNotFound      = errors.New("not Found")
	errAlreadyExists = errors.New("already Exists")
	errCantUpdate    = errors.New("cant't Update non-existing word")
)

func (d Dictionary) Search(word string) (string, error) {
	def, exists := d[word]

	if exists {
		return def, nil
	}

	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
		return nil
	}
	return errAlreadyExists
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
