package iterator

import "errors"

var ErrEndOfIterator = errors.New("Ene of Iterator")

type Iterator[T any] interface {
	HasNext() bool    // есть ли следующий (с текущего направления)
	HasPrev() bool    // есть ли предыдущий
	Next() (T, error) // следующий элемент
	Prev() (T, error) // предыдущий элемент
	Reset()           // сброс в начало
	Reverse()         // смена направления
}
