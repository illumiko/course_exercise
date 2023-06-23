package main

type Dictionary map[string]string
type Dictonary_err string

const (
	Err_not_found   = Dictonary_err("The word wasnt found")
	Err_word_exists = Dictonary_err("Word exists")
	Err_delete      = Dictonary_err("Words not found in dictionary cannot be deleted")
)

func (de Dictonary_err) Error() string {
	return string(de)
}

func (d Dictionary) Search(key string) (string, error) {
	word, found := d[key]
	if !found {
		return "", Err_not_found
	}
	return word, nil
}

func (d Dictionary) Add(key string, def string) (err error) {
	_, found := d.Search(key)
	switch found {
	case Err_not_found:
		d[key] = def
		return nil
	case nil:
		return Err_word_exists
	default:
		return found
	}
}

func (d Dictionary) Update(key string, def string) (err error) {
	_, found := d.Search(key)
	switch found {
	case nil:
		d[key] = def
		return nil
	case Err_not_found:
		return Err_not_found
	default:
		return found
	}
}

func (d Dictionary) Delete(key string) error {
	_, found := d.Search(key)
	if found != nil {
		return Err_delete
	}
	delete(d, key)
	return nil
}
