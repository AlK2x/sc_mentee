package iterator

type SliceIterator[T any] struct {
	reverse bool
	moved   bool
	idx     int
	s       []T
}

func NewSliceIterator[T any](s []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		reverse: false,
		idx:     0,
		s:       s,
	}
}

func (si *SliceIterator[T]) HasNext() bool {
	if si.reverse {
		return si.hasLeft()
	} else {
		return si.hasRight()
	}
}

func (si *SliceIterator[T]) HasPrev() bool {
	if si.reverse {
		return si.hasRight()
	} else {
		return si.hasLeft()
	}
}

func (si *SliceIterator[T]) Next() (T, error) {
	if si.reverse {
		return si.left()
	} else {
		return si.right()
	}
}

func (si *SliceIterator[T]) Prev() (T, error) {
	if si.reverse {
		return si.right()
	} else {
		return si.left()
	}
}

func (si *SliceIterator[T]) Reset() {
	si.moved = false
	si.idx = 0
}

func (si *SliceIterator[T]) Reverse() {
	if si.reverse {
		si.reverse = false
	} else {
		si.reverse = true
		if si.moved {
			si.idx--
		}
	}
}

func (si *SliceIterator[T]) hasLeft() bool {
	if len(si.s) == 0 {
		return false
	}
	return si.idx > 0
}

func (si *SliceIterator[T]) hasRight() bool {
	return si.idx < len(si.s)
}

func (si *SliceIterator[T]) left() (T, error) {
	if len(si.s) == 0 {
		var zero T
		return zero, ErrEndOfIterator
	}
	if si.idx == 0 {
		return si.s[si.idx], ErrEndOfIterator
	}
	si.idx--
	si.moved = true
	return si.s[si.idx], nil
}

func (si *SliceIterator[T]) right() (T, error) {
	if len(si.s) == 0 {
		var zero T
		return zero, ErrEndOfIterator
	}
	if si.idx == len(si.s) {
		return si.s[si.idx], ErrEndOfIterator
	}

	result := si.s[si.idx]
	si.idx++
	si.moved = true
	return result, nil
}
