package lib

type Slice[T any] []T
type Predicate[T any] func (iteratee T) bool

// Filter returns slice items matching specified predicate
func (slice Slice[T]) Filter(predicate Predicate[T]) Slice[T] {
	filtered := Slice[T]{}
	for _, v := range slice {
		if predicate == nil || predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// First returns the first item matching specified predicate
func (slice Slice[T]) First(predicate Predicate[T]) *T {
	for i, v := range slice {
		if predicate == nil || predicate(v) {
			return &slice[i]
		}
	}
	return nil
}

// Last returns the last item matching specified predicate
func (slice Slice[T]) Last(predicate Predicate[T]) *T {
	for i := len(slice)-1; i >= 0; i-- {
		if predicate == nil || predicate(slice[i]) {
			return &slice[i]
		}
	}
	return nil
}

// Any returns if any item matches the specified predicate
func (slice Slice[T]) Any(predicate Predicate[T]) bool {
	for _, v := range slice {
		if predicate == nil || predicate(v) {
			return true
		}
	}
	return false
}

// Every returns if every item matches the specified predicate
func (slice Slice[T]) Every(predicate Predicate[T]) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// Copy returns a copy of the slice
func (slice Slice[T]) Copy() Slice[T] {
	rSlice := make(Slice[T], len(slice))
	copy(rSlice, slice)
	return rSlice
}

// Reverse returns the slice in a reversed order
func (slice Slice[T]) Reverse() Slice[T] {
	rSlice := slice.Copy()
	for i, j := 0, len(rSlice)-1; i < j; i, j = i+1, j-1 {
		rSlice[i], rSlice[j] = rSlice[j], rSlice[i]
	}
	return rSlice
}