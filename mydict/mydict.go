package mydict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("not Found")

func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key]

	if exists {
		return value, nil
	}

	return "", errNotFound
}
