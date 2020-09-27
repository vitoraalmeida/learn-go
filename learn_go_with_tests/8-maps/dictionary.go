package main

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("connot update word because it doesn't exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

/*
	maps podem ser nil, então não se deve apenas declarar um map
	, pois o zero value é nil. Ler um map nil não tem problema
	mas escrever causa runtime panic

	var m map[string]string  --> causa runtime panic na escrita

	deve-se inicializar a variável com map vazio

	var dictionary = map[string]string{}
		ou
	var dictionary = make(map[string]string)

	isso cria um hash map vazio a aponta dictionary para ele.
*/

type Dictionary map[string]string

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

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
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
