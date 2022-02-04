package main

import "fmt"

type Slice[T any] []T
type Predicate[T any] func (iteratee T) bool

// Filter returns slice items matching specified predicate
func (slice Slice[T]) Filter(predicate Predicate[T]) []T {
	filtered := []T{}
	for _, v := range slice {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// First returns the first item matching specified predicate
func (slice Slice[T]) First(predicate Predicate[T]) *T {
	for i, v := range slice {
		if predicate(v) {
			return &slice[i]
		}
	}
	return nil
}

// Last returns the last item matching specified predicate
func (slice Slice[T]) Last(predicate Predicate[T]) *T {
	for i := len(slice)-1; i >= 0; i-- {
		if predicate(slice[i]) {
			return &slice[i]
		}
	}
	return nil
}

// Any returns if any item matches the specified predicate
func (slice Slice[T]) Any(predicate Predicate[T]) bool {
	for _, v := range slice {
		if predicate(v) {
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
func (slice Slice[T]) Copy() []T {
	rSlice := make(Slice[T], len(slice))
	copy(rSlice, slice)
	return rSlice
}

// Reverse returns the slice in a reversed order
func (slice Slice[T]) Reverse() []T {
	rSlice := slice.Copy()
	for i, j := 0, len(rSlice)-1; i < j; i, j = i+1, j-1 {
		rSlice[i], rSlice[j] = rSlice[j], rSlice[i]
	}
	return rSlice
}

type User struct {
	Name string
	Tickets int
}

func main() {
	slice := Slice[User]{
		{Name: "a", Tickets: 3},
		{Name: "b", Tickets: 1},
		{Name: "c", Tickets: 2},
	}
	fmt.Printf("Users with >= 2 tickets: %v\n", slice.Filter(func (u User) bool{
		return u.Tickets >= 2
	}))
	fmt.Printf("First with <= 2 tickets: %v\n", slice.First(func (u User) bool{
		return u.Tickets <= 2
	}))
	fmt.Printf("Last with >= 3 tickets: %v\n", slice.Last(func (u User) bool{
		return u.Tickets >= 3
	}))
	fmt.Printf("Last with >= 10 tickets: %v\n", slice.Last(func (u User) bool{
		return u.Tickets >= 10
	}))
	fmt.Printf("Has any with name \"foo\": %v\n", slice.Any(func (u User) bool{
		return u.Name == "foo"
	}))
	fmt.Printf("Every has > 0 tickets: %v\n", slice.Every(func (u User) bool{
		return u.Tickets > 0
	}))
	fmt.Printf("Slice address: %p / Slice copy address: %p\n", slice, slice.Copy())
	fmt.Printf("Slice in reverse: %v\n", slice.Reverse())
	fmt.Printf("Original Slice: %v\n", slice)
}