package iterable

type Iterable[T any] struct {
	next func() (T, bool, error)
}

func (i *Iterable[T]) Next() (T, bool, error) {
	return i.next()
}

func NewIterable[T any](next func() (T, bool, error)) *Iterable[T] {
	return &Iterable[T]{
		next: next,
	}
}
