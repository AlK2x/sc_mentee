package iterator

type FilterIterator[T any] struct {
	iter     Iterator[T]
	filterFn func(val T) bool
	buffer   T
	buffered bool
}

func NewFilterIterator[T any](iter Iterator[T], filterFn func(val T) bool) *FilterIterator[T] {
	return &FilterIterator[T]{
		iter:     iter,
		filterFn: filterFn,
	}
}

func (fi *FilterIterator[T]) HasNext() bool {
	if fi.buffered {
		return true
	}

	for !fi.buffered && fi.iter.HasNext() {
		val, _ := fi.iter.Next()
		if !fi.filterFn(val) {
			fi.buffered = true
			fi.buffer = val
		}
	}
	return fi.buffered
}

func (fi *FilterIterator[T]) HasPrev() bool {
	if fi.buffered {
		return true
	}

	for !fi.buffered && fi.iter.HasPrev() {
		val, _ := fi.iter.Prev()
		if !fi.filterFn(val) {
			fi.buffered = true
			fi.buffer = val
		}
	}
	return fi.buffered
}

func (fi *FilterIterator[T]) Next() (T, error) {
	if fi.buffered {
		fi.buffered = false
		return fi.buffer, nil
	}

	for !fi.buffered && fi.iter.HasNext() {
		val, _ := fi.iter.Next()
		if !fi.filterFn(val) {
			return val, nil
		}
	}
	var zero T
	return zero, ErrEndOfIterator

}

func (fi *FilterIterator[T]) Prev() (T, error) {
	if fi.buffered {
		fi.buffered = false
		return fi.buffer, nil
	}

	for !fi.buffered && fi.iter.HasPrev() {
		val, _ := fi.iter.Prev()
		if !fi.filterFn(val) {
			return val, nil
		}
	}
	var zero T
	return zero, ErrEndOfIterator
}

func (fi *FilterIterator[T]) Reset() {
	fi.iter.Reset()
	fi.buffered = false
}

func (fi *FilterIterator[T]) Reverse() {
	fi.iter.Reverse()
}
